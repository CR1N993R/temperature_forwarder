package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"temperature_forwarder/logs"
)

func GetRouter(clients []string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(418)
	})
	for _, client := range clients {
		log.Printf("Registering route /%s\n", client)
		router.Handle("/"+client, &logs.Controller{Context: client}).Methods(http.MethodPost)
	}
	return router
}
