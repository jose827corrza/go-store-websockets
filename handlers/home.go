package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/server"
)

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HomeResponse{
			Message: "Hello to the Go Store",
			Status:  true,
		})
	}
}
