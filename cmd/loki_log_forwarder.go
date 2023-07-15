package main

import (
	"fmt"
	"log"
	"loki-log-creator/config"
	"loki-log-creator/loki"
	"loki-log-creator/routes"
	"net/http"
)

func main() {
	configuration := config.LoadConfigFromFile()
	loki.CheckLoki(configuration.LokiInstances)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", configuration.Port),
		Handler: routes.GetRouter(configuration),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server failed to start")
	}
}
