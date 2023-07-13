package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var client http.Client

const url = "http://127.0.0.1:3100/loki/api/v1/push"

func main() {
	client = http.Client{}
	http.HandleFunc("/temperature", func(writer http.ResponseWriter, request *http.Request) {
		b, _ := io.ReadAll(request.Body)
		_ = request.Body.Close()
		logData(string(b))
	})
	log.Print("Server start on Port: 2112")
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		return
	}
}

func logData(data string) {
	logString := strings.ReplaceAll(data, "\"", "\\\"")
	timeStamp := int(time.Now().UnixNano())
	payload := strings.NewReader(`{"streams": [{"stream": {"home_automation": "temperature"},"values": [["` + strconv.Itoa(timeStamp) + `", "` + logString + `" ]]}]}`)
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	println(res.Status)
}
