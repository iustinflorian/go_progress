package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	wrappedHandler := loggingMiddleware(nextHandler)

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("loggingMiddleware returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if body := rr.Body.String(); body != "OK" {
		t.Errorf("loggingMiddleware altered response body: got %v want %v",
			body, "OK")
	}
}
