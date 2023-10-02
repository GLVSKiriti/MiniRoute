package handlers

import (
	"fmt"
	"net/http"
)

// URL shortening handlers
func (h *BaseHandler) Shorten(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello shorten", req.Context().Value("Uid"))
}
