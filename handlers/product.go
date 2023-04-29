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

func InsertProductHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dtos.ProductRequest

		// // Validation
		err := validators.ValidateCreateNewProduct(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		json.NewDecoder(r.Body).Decode(&request)
		// New uuid gen
		id := uuid.New()
		// hashedPsswrd, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		// Struct formation to be send to the repository
		var product = models.Product{
			Id: id.String(),
			// ID:          1,
			Name:        request.Name,
			Description: request.Description,
			Price:       request.Price,
			Stock:       request.Stock,
			Image:       request.Image,
			BrandID:     request.BrandID,
			CategoryID:  request.CategoryID,
		}

		//Insert using repository
		err = repository.InsertProduct(r.Context(), &product)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			utils.ErrorResponse(500, err.Error(), w)
			return
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
			CategoryID:  request.CategoryID,
		})
	}
}

func GetAllProductsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		products, err := repository.GetAllProducts(r.Context())
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(products)
	}
}

func UpdateAProductHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var request dtos.ProductRequest

		err := validators.ValidateCreateNewProduct(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		json.NewDecoder(r.Body).Decode(&request)

		var updatedProduct = models.Product{
			Id:          params["productId"],
			Name:        request.Name,
			Description: request.Description,
			Price:       request.Price,
			Stock:       request.Stock,
			Image:       request.Image,
			BrandID:     request.BrandID,
			CategoryID:  request.CategoryID,
		}

		product, err := repository.UpdateAProduct(r.Context(), params["productId"], &updatedProduct)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			utils.ErrorResponse(404, err.Error(), w)
			return
		}

		//Writting a successful insert
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}

func DeleteAProductHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := repository.DeleteAProduct(r.Context(), params["productId"])
		if err != nil {
			utils.ErrorResponse(404, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		test := map[string]string{"status": "OK", "message": "Product deleted successfully", "productId": params["productId"]}
		// utils.SuccessResponse(test,w)
		json.NewEncoder(w).Encode(test)
	}
}
