package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

type URLResult struct {
	URL        string `json:"url"`
	StatusCode int    `json:"status_code"`
	Title      string `json:"title"`
	Server     string `json:"server"`
}

func getTitle(html string) string {
	re := regexp.MustCompile("(?i)<title>(.*?)</title>")
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return "Title not found"
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logTime := time.Since(start)
			log.Printf("[%s] %s (Duration: %s)", r.Method, r.URL, logTime)
		})
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URLs []string `json:"urls"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	var results []URLResult
	ch := make(chan URLResult)
	client := &http.Client{Timeout: 5 * time.Second}

	for _, url := range body.URLs {
		go func(u string) {
			request, err := http.NewRequest("GET", u, nil)
			if err != nil {
				ch <- URLResult{URL: u, StatusCode: 0, Title: "URL Error"}
				return
			}
			request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

			response, err := client.Do(request)
			if err != nil {
				ch <- URLResult{URL: u, StatusCode: 0, Title: "URL error"}
				return
			}
			defer response.Body.Close()

			serverHeader := response.Header.Get("Server")

			bodyBytes, err := io.ReadAll(response.Body)
			titleText := "Error reading Body"
			if err == nil {
				titleText = getTitle(string(bodyBytes))
			}

			ch <- URLResult{
				URL:        u,
				StatusCode: response.StatusCode,
				Title:      titleText,
				Server:     serverHeader,
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
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
