package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
	"github.com/jose827corrza/go-store-websockets/utils"
	"github.com/jose827corrza/go-store-websockets/validators"
)

func InsertNewCategory(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dtos.CategoryRequest

		err := validators.ValidateCategory(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		json.NewDecoder(r.Body).Decode(&request)
		id := uuid.New()

		var category = models.Category{
			Id:    id.String(),
			Name:  request.Name,
			Image: request.Image,
		}

		err = repository.InsertCategory(r.Context(), &category)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.Category{
			Id:    id.String(),
			Name:  request.Name,
			Image: request.Image,
		})
	}
}

func GetProductsByCategory(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		log.Printf("path: %s", params["categoryId"])
		category, err := repository.GetAllProductsByCategory(r.Context(), params["categoryId"])
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(category)
	}
}

func GetAllCategories(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := repository.GetAllCategories(r.Context())
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(categories)
	}
}
