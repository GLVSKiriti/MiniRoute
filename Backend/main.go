package main

import (
	"log"
	"net/http"

	"github.com/GLVSKiriti/MiniRoute/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	authRouter := router.PathPrefix("/auth").Subrouter()
	urlRouter := router.PathPrefix("/url").Subrouter()

	routes.AuthRoutes(authRouter)
	routes.UrlRoutes(urlRouter)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
