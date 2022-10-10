package server

import (
	"net/http"
	"sesi10/server/controller"
)

var port = ":4000"

type NewRouter struct {
	controller con
}

func Start(port string) {
	router := http.NewServeMux()
	// endpoint := http.HandlerFunc(controller.GetUsers)
	router.HandleFunc("/users", Middleware1(controller.GetUsers()))

	http.ListenAndServe(port, router)

}
