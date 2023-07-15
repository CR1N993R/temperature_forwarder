package main

import (
	"log"
	"net/http"
	"temperature_forwarder/client"
	"temperature_forwarder/loki"
	"temperature_forwarder/routes"
)

func main() {
	if !loki.CheckLoki() {
		panic("Failed to start temperature forwarder unable to reach loki!")
	}
	clients := client.LoadConfigFromFile()
	server := &http.Server{
		Addr:    ":2113",
		Handler: routes.GetRouter(clients),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server failed to start")
	}
}
