package models

import "time"

// User model
type User struct {
	Uid      string `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// url model
type Url struct {
	Id        string    `json:"id"`
	Uid       string    `json:"uid"`
	LongUrl   string    `json:"longUrl"`
	ShortUrl  string    `json:"shortUrl"`
	CreatedOn time.Time `json:"createdOn"`
	ExpiresOn time.Time `json:"expiresOn"`
}
