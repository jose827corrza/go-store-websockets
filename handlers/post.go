package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

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

func CreateNewPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPost dtos.PostRequest

		err := validators.ValidatePost(&newPost, w, r)
		if err != nil {
			utils.ErrorResponse(400, err.Error(), w)
			return
		}

		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		userInfo, err, status := utils.GetUserInfoByJWTToken(tokenString, s, r)
		if err != nil {
			utils.ErrorResponse(status, err.Error(), w)
			return
		}

		id := uuid.New()

		var post = models.Post{
			ID:          id.String(),
			Title:       newPost.Title,
			Description: newPost.Description,
			UserID:      userInfo.Id,
		}

		err = repository.CreatePost(r.Context(), &post)
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	}
}

func CreateUserPostHandler(s server.Server) http.HandlerFunc {
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
			Role:     "postmaker",
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

func GetAllPostsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := repository.GetAllPosts(r.Context())
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(posts)
	}
}

func GetAllPostsByUserIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParameter := mux.Vars(r)

		userPosts, err := repository.GetAllPostsByUserId(r.Context(), pathParameter["userId"])
		if err != nil {
			utils.ErrorResponse(500, err.Error(), w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userPosts)
	}
}
