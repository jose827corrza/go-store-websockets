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

func UserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dtos.SignUpUserRequest

		// Validation
		err := validators.ValidateSignUp(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}

		// New uuid gen
		id := uuid.New()
		hashedPsswrd, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Struct formation to be send to the repository
		var user = models.User{
			Id:       id.String(),
			Email:    request.Email,
			Password: string(hashedPsswrd),
		}

		//Insert using repository
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//Writting a successful insert
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(dtos.SignUpUserResponse{
			Id:    id.String(),
			Email: user.Email,
		})
	}
}
