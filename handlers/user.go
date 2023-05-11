package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
	"github.com/jose827corrza/go-store-websockets/utils"
	"github.com/jose827corrza/go-store-websockets/validators"
	"golang.org/x/crypto/bcrypt"
)

func UserHandlerInsert(s server.Server) http.HandlerFunc {
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
			Role:     "administrator",
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

func UserHandlerGetById(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		user, err := repository.GetUserById(r.Context(), params["userId"])
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func UserHandlerGetAll(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.GetAllUsers(r.Context())
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request dtos.SignUpUserRequest

		// Validation
		err := validators.ValidateSignUp(&request, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}

		// Lookfor into the DB the user with that email
		user, err := repository.GetUserByEmail(r.Context(), request.Email)
		if err != nil {
			utils.ErrorResponse(404, err.Error(), w)
			return
		}

		// Compare both psswd
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			// If does not match returns a 400
			utils.ErrorResponse(400, "Wrong credentials", w)
			return
		}

		// From here, we can assume
		// - Exists that user
		// - Has the right credentials
		claims := models.AppClaims{
			UserId: user.Id,
			Role:   user.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		// For send info as a cookie
		// cookie := http.Cookie{
		// 	Name:       "test",
		// 	RawExpires: time.Now().Add(2 * time.Hour).String(),
		// 	Value:      tokenString,
		// }
		// http.SetCookie(w, &cookie)
		s.Hub().BroadCast(claims.UserId, nil)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dtos.LoginResponse{
			Token: tokenString,
		})
		// r.Cookie("test")
	}
}

func MeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		userInfo, err, statusInt := utils.GetUserInfoByJWTToken(tokenString, s, r)
		if err != nil {
			utils.ErrorResponse(statusInt, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusInt)
		json.NewEncoder(w).Encode(userInfo)
	}
}
