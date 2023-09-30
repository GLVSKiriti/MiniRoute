package routes

import (
	"github.com/GLVSKiriti/MiniRoute/handlers"
	"github.com/gorilla/mux"
)

// Routes Here
func AuthRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/login", handlers.Login).Methods("GET")
}

func UrlRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/shorten", handlers.Shorten).Methods("GET")
}
