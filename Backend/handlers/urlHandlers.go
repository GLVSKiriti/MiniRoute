package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GLVSKiriti/MiniRoute/models"
)

// URL shortening handlers
func (h *BaseHandler) Shorten(res http.ResponseWriter, req *http.Request) {
	url := &models.Url{}
	json.NewDecoder(req.Body).Decode(url)
	uid := int(req.Context().Value("Uid").(float64))

	var maxUrlId int

	err := h.db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM urlmappings WHERE uid = $1", uid).Scan(&maxUrlId)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Increment the url id
	nextUrlId := maxUrlId + 1

	// Url short code
	shortCode := fmt.Sprintf("%d-%d", uid, nextUrlId)
	// Insert url in database

	_, err2 := h.db.Exec("INSERT INTO urlmappings (uid,id,longurl,shorturl) VALUES ($1,$2,$3,$4)", uid, nextUrlId, url.LongUrl, shortCode)

	if err2 != nil {
		http.Error(res, err2.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{"shortUrl": shortCode}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}
