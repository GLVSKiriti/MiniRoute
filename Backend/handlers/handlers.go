package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GLVSKiriti/MiniRoute/models"
)

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{db: db}
}

// Authentication Handlers
func (h *BaseHandler) Login(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	// Check whether user exists or not
	var password string
	err := h.db.QueryRow(`SELECT password FROM users WHERE email=$1`, user.Email).Scan(&password)
	if err == sql.ErrNoRows {
		http.Error(res, "User not Found! Register Instead", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check password if user exists
	if password != user.Password {
		http.Error(res, "Wrong Password!!", http.StatusUnauthorized)
		return
	}

	//Password is correct
	res.WriteHeader(http.StatusOK)
}

func (h *BaseHandler) Register(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	// Check whether email exists or not
	var emailExists int
	err := h.db.QueryRow(`SELECT COUNT(*) FROM users WHERE email=$1;`, user.Email).Scan(&emailExists)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if emailExists > 0 {
		http.Error(res, "User Alreaday exists", http.StatusConflict)
		return
	}

	// Save the details in database
	_, err = h.db.Exec(`INSERT INTO users (email,password) VALUES($1,$2);`, user.Email, user.Password)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
}

// URL shortening handlers
func (h *BaseHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello shorten")
}
