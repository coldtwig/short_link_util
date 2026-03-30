package middleware

import (
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
