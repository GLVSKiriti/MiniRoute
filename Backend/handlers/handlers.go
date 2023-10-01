package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{db: db}
}

// Authentication Handlers
func (h *BaseHandler) Login(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Hello Man")

	sqlStatement := `INSERT INTO users (email,password) VALUES('ABC@GMAIL.COM','123456');`
	_, err := h.db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func (h *BaseHandler) Register(res http.ResponseWriter, req *http.Request) {

}

// URL shortening handlers
func (h *BaseHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello shorten")
}
