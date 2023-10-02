package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/GLVSKiriti/MiniRoute/models"
	"github.com/golang-jwt/jwt/v5"
)

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{db: db}
}

// Login Handler
func (h *BaseHandler) Login(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	// Check whether user exists or not
	var password string
	var uid int
	err := h.db.QueryRow(`SELECT uid,password FROM users WHERE email=$1`, user.Email).Scan(&uid, &password)
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
	// Generate a JWT and send it
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Uid"] = uid
	secret_key := []byte(os.Getenv("SECRETKEY"))
	tokenStr, err := token.SignedString(secret_key)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{"Authorization": tokenStr}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}

// Register Handler
func (h *BaseHandler) Register(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	// Check whether email exists or not
	var emailExists int
	var uid int
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
	err = h.db.QueryRow(`INSERT INTO users (email,password) VALUES($1,$2) RETURNING uid;`, user.Email, user.Password).Scan(&uid)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Uid"] = uid
	secret_key := []byte(os.Getenv("SECRETKEY"))
	tokenStr, err := token.SignedString(secret_key)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{"Authorization": tokenStr}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(data)

}
