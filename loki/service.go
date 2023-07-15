package loki

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CheckLoki(lokiInstances []string) {
	for _, instance := range lokiInstances {
		req, _ := http.NewRequest("GET", instance+"/ready", nil)
		req.Header.Add("Content-Type", "application/json")

		httpClient := http.Client{}
		res, err := httpClient.Do(req)
		if err != nil {
			log.Println(err)
			panic("Failed to start temperature forwarder unable to reach loki instance: " + instance + "!")
		}
		if res.StatusCode != http.StatusOK {
			panic("Failed to start temperature forwarder unable to reach loki instance: " + instance + "!")
		}
	}
}

func SendLogToLoki(logString string, client string, lokiInstances []string) error {
	logString = strings.ReplaceAll(logString, "\"", "\\\"")
	logString = strings.ReplaceAll(logString, "\n", "")
	timeStamp := int(time.Now().UnixNano())
	payload := strings.NewReader(`{"streams": [{"stream": {"temperature_forwarder": "` + client + `"},"values": [["` + strconv.Itoa(timeStamp) + `", "` + logString + `" ]]}]}`)
	for _, instance := range lokiInstances {
		req, err := http.NewRequest("POST", instance+"/loki/api/v1/push", payload)
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
		if res.StatusCode != http.StatusNoContent {
			defer res.Body.Close()
			log.Println(res.Status)
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println(string(resBody))
		}
	}
	return nil
}
