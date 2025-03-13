package middleware

import (
	"net/http"
	"os"

	"github.com/bootcamp-go/web/response"
)

func Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("API_TOKEN") {
			response.Error(w, http.StatusUnauthorized, "Token inválido.")
			return
		}
		handler.ServeHTTP(w, r)
	})
}
