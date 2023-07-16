package loki

import (
	"io"
	"log"
	"loki-log-creator/config"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SendLogToLoki(logMessage string, client string, lokiInstances []config.LokiInstance) error {
	for _, instance := range lokiInstances {
		url := instance.URL + "/loki/api/v1/push"
		req, err := http.NewRequest(http.MethodPost, url, getPayload(client, logMessage))
		if err != nil {
			log.Println(err)
			return err
		}
		req.Header.Add("Content-Type", "application/json")
		if instance.Token != "" {
			req.Header.Add("Authorization", "Bearer "+instance.Token)
		}

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

func getPayload(client string, log string) io.Reader {
	log = strings.ReplaceAll(log, "\"", "\\\"")
	log = strings.ReplaceAll(log, "\n", "")
	timeStamp := int(time.Now().UnixNano())
	return strings.NewReader(`{"streams": [{"stream": {"loki_log_forwarder": "` + client + `"},"values": [["` + strconv.Itoa(timeStamp) + `", "` + log + `" ]]}]}`)
}
