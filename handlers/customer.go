package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
	"github.com/jose827corrza/go-store-websockets/utils"
	"github.com/jose827corrza/go-store-websockets/validators"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewCustomerHanlder(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dtos.SignUpCustomerRequest

		err := validators.ValidateCustomerSignUp(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}
		id := uuid.New()
		userId := uuid.New()
		hashedPsswrd, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), 10)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var requestCustomer = models.User{
			Id:       userId.String(),
			Email:    request.User.Email,
			Password: string(hashedPsswrd),
			Role:     "customer",
		}
		var customer = models.Customer{
			Id:       id.String(),
			Email:    request.Email,
			Name:     request.Name,
			LastName: request.LastName,
			Phone:    request.Phone,
			User:     requestCustomer,
		}

		//Informs in websocket the new customer creation
		s.Hub().BroadCast(requestCustomer, nil)
		err = repository.InsertUser(r.Context(), &requestCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = repository.InsertCustomer(r.Context(), &customer)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(dtos.SignUpUserResponse{
			Id:    id.String(),
			Email: customer.User.Email,
		})
	}
}
