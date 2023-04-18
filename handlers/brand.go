package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
)

func InsertNewBrand(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.Brand

		json.NewDecoder(r.Body).Decode(&request)
		id := uuid.New()

		var brand = models.Brand{
			Id:    id.String(),
			Name:  request.Name,
			Image: request.Image,
		}

		err := repository.InsertBrand(r.Context(), &brand)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.Brand{
			Id:    id.String(),
			Name:  request.Name,
			Image: request.Image,
		})
	}
}
