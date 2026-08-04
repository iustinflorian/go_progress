package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logTime := time.Since(start)
			log.Printf("[%s] %s (Duration: %s)", r.Method, r.URL, logTime)
		})
}
