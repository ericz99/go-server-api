package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Log Struct (Model)
type Log struct {
	RequestURI string
	Method     string
}

// LoggerMiddleware METHOD
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logMSG := Log{RequestURI: r.RequestURI, Method: r.Method}
		// # print log message
		log.Printf("[%s] [%s] [%s]", logMSG.Method, logMSG.RequestURI, time.Now().Format(time.RFC850))
		// # forward to next middleware
		next.ServeHTTP(w, r)
	})
}
