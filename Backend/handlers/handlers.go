package handlers

import (
	"fmt"
	"net/http"
)

// handlers here
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello Man")
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello shorten")
}
