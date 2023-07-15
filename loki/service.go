package loki

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const url = "http://192.168.1.16:3100"

func CheckLoki() bool {
	req, err := http.NewRequest("GET", url+"/ready", nil)
	if err != nil {
		log.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	return res.StatusCode == http.StatusOK
}

func SendLogToLoki(logString string, client string) error {
	logString = strings.ReplaceAll(logString, "\"", "\\\"")
	logString = strings.ReplaceAll(logString, "\n", "")
	timeStamp := int(time.Now().UnixNano())
	payload := strings.NewReader(`{"streams": [{"stream": {"temperature_forwarder": "` + client + `"},"values": [["` + strconv.Itoa(timeStamp) + `", "` + logString + `" ]]}]}`)
	req, err := http.NewRequest("POST", url+"/loki/api/v1/push", payload)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		log.Println(res.Status)
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(string(resBody))
	}
	return nil
}
