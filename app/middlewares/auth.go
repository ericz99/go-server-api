package middlewares

import (
	"context"
	"encoding/json"
	Auth "go-server-api/app/auth"
	"net/http"
	"strings"
)

// Exception struct (Model)
type Exception struct {
	Message string `json:"message"`
}

// AuthMiddleware METHOD
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// # check if auth token is visible in header
		if r.Header.Get("x-auth-token") == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing header token"})
			return
		}

		header := strings.TrimSpace(r.Header.Get("x-auth-token"))
		// # decode auth token
		tk, err := Auth.DecodeToken(header)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		// # forward next middlware
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
