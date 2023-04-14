package middlewares

import (
	"log"
	"net/http"
)

func LogginMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		log.Println(r.Header.Values("User-Agent"))
		f(w, r) // Important to return the values and let pass the other middlewares or handlers bihind
	}
}
