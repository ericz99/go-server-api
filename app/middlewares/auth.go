package middlewares

import "net/http"

// AuthMiddleware METHOD
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// # check if auth token is visible in header
		if r.Header.Get("x-auth-token") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// # forward next middlware
		next.ServeHTTP(w, r)
	})
}
