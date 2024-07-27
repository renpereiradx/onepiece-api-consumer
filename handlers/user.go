package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/renpereiradx/onepiece-api-consumer/models"
	"github.com/renpereiradx/onepiece-api-consumer/repository"
	"github.com/renpereiradx/onepiece-api-consumer/server"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_COST = 8
)

type SignupResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func SignupHandler(server server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := SignupRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASH_COST)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := models.User{
			ID:       id.String(),
			Password: string(hashedPassword),
			Email:    request.Email,
		}
		err = repository.InsertUser(context.Background(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignupResponse{
			ID:    user.ID,
			Email: user.Email,
		})
	}
}
