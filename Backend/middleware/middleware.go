package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// middlewares here
func VerifyToken(handler http.HandlerFunc) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("Authorization")
		var claims jwt.MapClaims
		// Below func first parses and extract claims in token to claims parameter
		// Third argument, which is a callback function. This function is responsible for validating the token.
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRETKEY")), nil
		})
		if err != nil || !token.Valid {
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(map[string]string{"Message": err.Error()})
			return
		}
		//The below line is indirectly just runs the passed handler by setting uid in context
		// so handlers can use the uid
		ctx := context.WithValue(req.Context(), "Uid", claims["Uid"])
		handler.ServeHTTP(res, req.WithContext(ctx))
	}
}
