package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type URLResult struct {
	URL        string `json:"url"`
	StatusCode int    `json:"status_code"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logTime := time.Since(start)
			log.Printf("[%s] %s (duration: %s)", r.Method, r.URL, logTime)
		})
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URLs []string `json:"urls"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	var results []URLResult

	ch := make(chan URLResult)

	for _, url := range body.URLs {
		go func(u string) {
			response, err := http.Get(u)
			if err == nil {
				ch <- URLResult{URL: u, StatusCode: response.StatusCode}
				response.Body.Close()
			} else {
				ch <- URLResult{URL: u, StatusCode: 0}
			}
		}(url)
	}

	for i := 0; i < len(body.URLs); i++ {
		res := <-ch
		results = append(results, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	wrappedHandler := loggingMiddleware(http.HandlerFunc(handler))
	http.Handle("POST /check", wrappedHandler)
	http.ListenAndServe(":8080", nil)
}
