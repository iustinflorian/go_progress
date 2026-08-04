package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()

	http.Handle("POST /check", loggingMiddleware(http.HandlerFunc(postRequestHandler)))
	http.Handle("GET /sessions/", loggingMiddleware(http.HandlerFunc(getSessionsHandler)))
	http.Handle("GET /sessions/{id}", loggingMiddleware(http.HandlerFunc(getSessionByIDHandler)))
	http.Handle("DELETE /sessions/rm/{id}", loggingMiddleware(http.HandlerFunc(deleteSessionHandler)))

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
