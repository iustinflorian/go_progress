package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTitle(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "Valid title tag",
			html:     "<html><head><title>Google</title></head><body></body></html>",
			expected: "Google",
		},
		{
			name:     "Uppercase TITLE tag",
			html:     "<html><head><TITLE>GitHub · Change is constant</TITLE></head></html>",
			expected: "GitHub · Change is constant",
		},
		{
			name:     "Missing title tag",
			html:     "<html><head></head><body>Page without title</body></html>",
			expected: "Title not found",
		},
		{
			name:     "Empty HTML",
			html:     "",
			expected: "Title not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTitle(tt.html)
			if got != tt.expected {
				t.Errorf("getTitle() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestPostRequestHandler_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><head><title>Test Page</title></head></html>"))
	}))
	defer mockServer.Close()

	jsonBody := []byte(`{"urls": ["` + mockServer.URL + `"]}`)
	req, err := http.NewRequest("POST", "/check", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postRequestHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("postRequestHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostRequestHandler_InvalidJSONStruct(t *testing.T) {
	invalidJSON := []byte(`{"urls": [invalid-json}`)

	req, err := http.NewRequest("POST", "/check", bytes.NewBuffer(invalidJSON))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postRequestHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("postRequestHandler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetSessionByIDHandler_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/sessions/123-invalid-id", nil)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	req.SetPathValue("id", "123-invalid-id")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSessionByIDHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("getSessionByIDHandler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
