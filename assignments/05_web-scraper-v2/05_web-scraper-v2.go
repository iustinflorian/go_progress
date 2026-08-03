package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CheckSession struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	TotalURLs int           `bson:"total_urls" json:"total_urls"`
	Results   []URLResult   `bson:"results" json:"results"`
}

type URLResult struct {
	URL        string `json:"url" bson:"url"`
	StatusCode int    `json:"status_code" bson:"status_code"`
	Title      string `json:"title,omitempty" bson:"title,omitempty"`
	Server     string `json:"server,omitempty" bson:"server,omitempty"`
}

var sessionCollection *mongo.Collection

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

	session := CheckSession{
		CreatedAt: time.Now(),
		TotalURLs: len(results),
		Results: results,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertResult, err := sessionCollection.InsertOne(ctx, session)
	if err != nil {
		log.Fatalf("Error saving to MongoDB: %v", err)
	} else {
		log.Printf("Session successfully saved in MongoDB with ID: %v", insertResult.InsertedID)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	mongoURI := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB doesn't respond: %v", err)
	}
	log.Println("MongoDB connection established")

	sessionCollection = client.Database("web-scraper_db").Collection("sessions")

	wrappedHandler := loggingMiddleware(http.HandlerFunc(handler))
	http.Handle("POST /check", wrappedHandler)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
