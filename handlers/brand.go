package handlers

import (
	"encoding/json"
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

func InsertNewBrand(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dtos.BrandRequest

		// json.NewDecoder(r.Body).Decode(&request)
		err := validators.ValidateBrand(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		id := uuid.New()

		var brand = models.Brand{
			Id:    id.String(),
			Name:  request.Name,
			Image: request.Image,
		}

		err = repository.InsertBrand(r.Context(), &brand)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			utils.ErrorResponse(500, err.Error(), w)
			return
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

func GetAllProductsByBrandHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		brand, err := repository.GetAllProductsByBrand(r.Context(), params["brandId"])
		if err != nil {
			utils.ErrorResponse(404, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&brand)
	}
}
