package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func LogginMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		log.Println(r.Header.Values("User-Agent"))
		f(w, r) // Important to return the values and let pass the other middlewares or handlers bihind
	}
}

var validate *validator.Validate

func SignUpValidator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		f(w, r)
	}
}

type SignUpDTO struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
