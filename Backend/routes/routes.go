package routes

import (
	"github.com/GLVSKiriti/MiniRoute/handlers"
	"github.com/GLVSKiriti/MiniRoute/middleware"
	"github.com/gorilla/mux"
)

// Authentication Routes
func AuthRoutes(subRouter *mux.Router, h *handlers.BaseHandler) {
	subRouter.HandleFunc("/login", h.Login).Methods("POST")
	subRouter.HandleFunc("/register", h.Register).Methods("POST")
}

// URL shortneing Routes
func UrlRoutes(subRouter *mux.Router, h *handlers.BaseHandler) {
	subRouter.HandleFunc("/shorten", middleware.VerifyToken(h.Shorten)).Methods("POST")
	subRouter.HandleFunc("/redirect/{shortCode}", h.RedirectToOriginalUrl)

	subRouter.HandleFunc("/myurls", middleware.VerifyToken((h.GetMyUrls))).Methods("GET")
}
