package helpers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/renpereiradx/onepiece-api-consumer/models"
	"github.com/renpereiradx/onepiece-api-consumer/server"
)

func GetJWTAuthorizationInfo(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, models.Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	return token, err
}
