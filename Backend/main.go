package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GLVSKiriti/MiniRoute/config"
	"github.com/GLVSKiriti/MiniRoute/handlers"
	"github.com/GLVSKiriti/MiniRoute/routes"
	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	header := muxHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := muxHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := muxHandlers.AllowedOrigins([]string{"*"})

	authRouter := router.PathPrefix("/auth").Subrouter()
	urlRouter := router.PathPrefix("/url").Subrouter()

	// Initializing the database
	db := config.InitDb()
	fmt.Printf("Database Connected")
	defer db.Close()

	h := handlers.NewBaseHandler(db)

	routes.AuthRoutes(authRouter, h)
	routes.UrlRoutes(urlRouter, h)

	err := http.ListenAndServe(":8080", muxHandlers.CORS(header, methods, origins)(router))
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
