package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/renpereiradx/onepiece-api-consumer/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(server server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome To One Piece API",
			Status:  true,
		})
	}
}
