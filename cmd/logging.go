package main

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before the handler
		startTime := time.Now()

		// After the handler
		log.Printf("%s %s %s %v", r.RemoteAddr, r.Method, r.URL.Path, time.Since(startTime))

		next.ServeHTTP(w, r) // Call the next handler, which can be another middleware in the chain or the final handler.
	})
}
