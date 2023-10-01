package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GLVSKiriti/MiniRoute/config"
	"github.com/GLVSKiriti/MiniRoute/handlers"
	"github.com/GLVSKiriti/MiniRoute/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	authRouter := router.PathPrefix("/auth").Subrouter()
	urlRouter := router.PathPrefix("/url").Subrouter()

	// Initializing the database
	db := config.InitDb()
	fmt.Printf("Database Connected")
	defer db.Close()

	h := handlers.NewBaseHandler(db)

	routes.AuthRoutes(authRouter, h)
	routes.UrlRoutes(urlRouter, h)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
