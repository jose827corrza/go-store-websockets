package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jose827corrza/go-store-websockets/models"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/server"
	"github.com/jose827corrza/go-store-websockets/utils"
)

var (
	//Path that dont require the auth validation
	NO_AUTH_NEED = []string{
		"login",
		"signup",
		"categories",
		"brands",
		"products",
		"users",
		"customers",
		"posts",
		"",
	}
)

func shouldCheckToken(path string, method string) bool {
	for _, p := range NO_AUTH_NEED {
		if strings.Contains(path, p) {
			if method == "PUT" {
				return true
			}
			if method == "POST" && strings.Contains(path, "create") {
				return true
			}
			return false
		}
	}
	return true
}

func AdminPriviledges(method string, role string) bool {
	if method == "POST" || method == "PUT" {
		if role != "administrator" {
			return false
		}
		return true
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path, r.Method) {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				utils.ErrorResponse(401, err.Error(), w)
				// http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
				_, err := repository.GetUserById(r.Context(), claims.UserId)
				if err != nil {
					utils.ErrorResponse(500, err.Error(), w)
					return
				}

			}
			next.ServeHTTP(w, r)
		})
	}
}
