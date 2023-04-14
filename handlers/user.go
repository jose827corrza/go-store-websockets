package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
	"golang.org/x/crypto/bcrypt"
)

// NORMAL
// type SignUpUserRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// USING VALIDATE
type SignUpUserRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type SignUpUserResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func UserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request SignUpUserRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		id := uuid.New()
		hashedPsswrd, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var user = models.User{
			Id:       id.String(),
			Email:    request.Email,
			Password: string(hashedPsswrd),
			// Password: request.Password,
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SignUpUserResponse{
			Id:    id.String(),
			Email: user.Email,
		})
	}
}
