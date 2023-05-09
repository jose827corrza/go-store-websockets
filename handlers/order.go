package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
	"github.com/jose827corrza/go-store-websockets/utils"
	"github.com/jose827corrza/go-store-websockets/validators"
)

func CreateAnOrderHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dtos.NewOrderRequest

		err := validators.ValidateNewOrder(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		id := uuid.New()

		var newOrder = models.Order{
			Id:         id.String(),
			CustomerID: request.CustomerID,
			Items:      request.Items,
			Date:       time.Now(),
		}

		err = repository.CreateAnOrder(r.Context(), &newOrder)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.Order{
			Id:         newOrder.Id,
			CustomerID: newOrder.CustomerID,
			Items:      newOrder.Items,
		})
	}
}

func GetOrdersByCustomerIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		customer, err := repository.GetOrdersByCustomerId(r.Context(), params["customerId"])
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customer)
	}
}

func AddAProductToAnOrderHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.OrderItem

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		id := uuid.New()
		err = repository.AddAProductToAnOrder(r.Context(), request.OrderID, request.ProductID, id.String())
		if err != nil {
			utils.ErrorResponse(500, "Error adding the product bro :c", w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(request)
	}
}

func GetOrderByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		order, err := repository.GetOrderById(r.Context(), params["orderId"])
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}
}
