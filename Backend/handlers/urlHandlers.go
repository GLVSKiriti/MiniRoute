package handlers

import (
	"fmt"
	"net/http"
)

// URL shortening handlers
func (h *BaseHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello shorten")
}
