package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
)

func SuccessResponse(fields map[string]interface{}, writer http.ResponseWriter) {
	fields["status"] = "success"
	message, err := json.Marshal(fields)
	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}

func ErrorResponse(statusCode int, error string, writer http.ResponseWriter) {
	//Create a new map and fill it
	fields := make(map[string]interface{})
	fields["status"] = fmt.Sprintf("%d", statusCode)
	fields["message"] = error
	message, err := json.Marshal(fields)

	if err != nil {
		//An error occurred processing the json
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
		return
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(message)
}

func GetUserInfoByJWTToken(tokenString string, s server.Server, r *http.Request) (*models.User, error, int) {
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		// utils.ErrorResponse(401, err.Error(), w)
		// http.Error(w, err.Error(), http.StatusUnauthorized)
		return nil, err, 401
	}
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		user, err := repository.GetUserById(r.Context(), claims.UserId)
		if err != nil {
			// utils.ErrorResponse(500, err.Error(), w)
			return nil, err, 500
		}

		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(user)
		return user, nil, 200
	} else {
		// utils.ErrorResponse(500, err.Error(), w)
		return nil, err, 500
	}
}
