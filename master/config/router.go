package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateRouter CreateRouter
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

//RunServer RunServer
func RunServer(router *mux.Router) {
	server := GetCustomConf("server", "default")
	port := GetCustomConf("port", "default")
	serverAndPort := fmt.Sprintf("%v:%v", server, port)
	fmt.Println("Setting Web Server at port : " + server + ":" + port)
	err := http.ListenAndServe(serverAndPort, router)
	if err != nil {
		log.Fatal(err)
	}
}
