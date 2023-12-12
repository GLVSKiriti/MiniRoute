package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GLVSKiriti/MiniRoute/models"
	"github.com/gorilla/mux"
)

// URL shortening handlers

func (h *BaseHandler) Shorten(res http.ResponseWriter, req *http.Request) {
	url := &models.Url{}
	json.NewDecoder(req.Body).Decode(url)
	uid := int(req.Context().Value("Uid").(float64))
	var shortCode string

	if url.CustomShortUrl != nil {
		var count int
		err := h.db.QueryRow("SELECT COUNT(*) FROM urlmappings WHERE shorturl = $1", url.CustomShortUrl).Scan(&count)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if count > 0 {
			http.Error(res, "ShortCode already exists", http.StatusConflict)
			return
		}

		shortCode = *url.CustomShortUrl
	}

	var maxUrlId int
	err := h.db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM urlmappings WHERE uid=$1", uid).Scan(&maxUrlId)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Next id as we are inserting new row
	maxUrlId = maxUrlId + 1

	if url.CustomShortUrl == nil {
		shortCode = fmt.Sprintf("%d-%d", uid, maxUrlId)
	}

	_, err = h.db.Exec("INSERT INTO urlmappings (uid,id,longurl,shorturl) VALUES ($1,$2,$3,$4)", uid, maxUrlId, url.LongUrl, shortCode)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{"shortUrl": shortCode}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}

func (h *BaseHandler) RedirectToOriginalUrl(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	shortCode := vars["shortCode"]

	var longUrl string

	err := h.db.QueryRow("SELECT longurl FROM urlmappings WHERE shorturl = $1", shortCode).Scan(&longUrl)

	if err != nil {
		http.Error(res, "URL not found", http.StatusNotFound)
		return
	}

	// Redierct to original url
	http.Redirect(res, req, longUrl, http.StatusSeeOther)
}

func (h *BaseHandler) GetMyUrls(res http.ResponseWriter, req *http.Request) {
	uid := int(req.Context().Value("Uid").(float64))

	rows, err := h.db.Query("SELECT * FROM urlmappings WHERE uid=$1", uid)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var urls []models.Url
	for rows.Next() {
		var url models.Url
		err := rows.Scan(&url.Uid, &url.Id, &url.LongUrl, &url.CustomShortUrl)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(urls)
}

//rows.Next() indicates whether the next row from result set is available, and will return true until either result set is exhausted or an error has occurred during fetching the data. For this reason you should always check for an error at the end of the for rows.Next() loop (this is done calling rows.Err()). If there’s an error during the loop, you need to know about it. Don’t just assume that the loop iterates until you’ve processed all the rows.
