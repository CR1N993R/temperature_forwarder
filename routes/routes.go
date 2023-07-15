package routes

import (
	"github.com/gorilla/mux"
	"log"
	"loki-log-creator/config"
	"loki-log-creator/logs"
	"net/http"
)

func GetRouter(configuration config.Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(418)
	})
	for _, client := range configuration.Clients {
		log.Printf("Registering route /%s\n", client)
		router.Handle("/"+client, &logs.Controller{Context: client}).Methods(http.MethodPost)
	}
	return router
}
