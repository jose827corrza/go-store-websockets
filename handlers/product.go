package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
)

func InsertProductHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.Product

		// // Validation
		// err := validators.ValidateSignUp(&request, w, r)
		// if err != nil {
		// 	utils.ErrorResponse(400, err.Error(), w)
		// 	return
		// }
		json.NewDecoder(r.Body).Decode(&request)
		// New uuid gen
		id := uuid.New()
		// hashedPsswrd, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		// Struct formation to be send to the repository
		var product = models.Product{
			Id:          id.String(),
			Name:        request.Name,
			Description: request.Description,
			Price:       request.Price,
			Stock:       request.Stock,
			Image:       request.Image,
			BrandID:     request.BrandID,
		}

		//Insert using repository
		err := repository.InsertProduct(r.Context(), &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//Writting a successful insert
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.Product{
			Id:          id.String(),
			Name:        request.Name,
			Description: request.Description,
			Price:       request.Price,
			Stock:       request.Stock,
			Image:       request.Image,
			BrandID:     request.BrandID,
		})
	}
}
