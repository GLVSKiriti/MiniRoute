package models

// User model
type User struct {
	Uid      int    `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// url model
type Url struct {
	Uid            string `json:"uid"`
	LongUrl        string `json:"longUrl"`
	CustomShortUrl string `json:"shortUrl"`
}
