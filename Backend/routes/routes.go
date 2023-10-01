package routes

import (
	"github.com/GLVSKiriti/MiniRoute/handlers"
	"github.com/gorilla/mux"
)

// Authentication Routes
func AuthRoutes(subRouter *mux.Router, h *handlers.BaseHandler) {
	subRouter.HandleFunc("/login", h.Login).Methods("GET")
	subRouter.HandleFunc("/register", h.Register).Methods("POST")
}

// URL shortneing Routes
func UrlRoutes(subRouter *mux.Router, h *handlers.BaseHandler) {
	subRouter.HandleFunc("/shorten", h.Shorten).Methods("GET")
}
