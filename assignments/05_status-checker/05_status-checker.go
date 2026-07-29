package main

import (
	"encoding/json"
	"net/http"
)

type URLResult struct {
	URL        string
	StatusCode int
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URLs []string
	}

	json.NewDecoder(r.Body).Decode(&body)

	var results []URLResult

	for _, url := range body.URLs {
		response, err := http.Get(url)
		if err == nil {
			results = append(results, URLResult{URL: url, StatusCode: response.StatusCode})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("POST /check", handler)
	http.ListenAndServe(":8080", nil)
}
