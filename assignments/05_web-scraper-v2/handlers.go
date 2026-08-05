package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getTitle(html string) string {
	re := regexp.MustCompile("(?i)<title>(.*?)</title>")
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return "Title not found"
}

func postRequestHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URLs []string `json:"urls"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
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
		Results:   results,
	}

	if sessionCollection != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		insertResult, err := sessionCollection.InsertOne(ctx, session)
		if err != nil {
			log.Printf("Error saving to MongoDB: %v", err)
		} else {
			log.Printf("Session saved with ID: %v", insertResult.InsertedID)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func getSessionsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := sessionCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch sessions from DB", http.StatusInternalServerError)
		log.Printf("Error finding sessions: %v", err)
		return
	}
	defer cursor.Close(ctx)

	var sessions []CheckSession
	if err := cursor.All(ctx, &sessions); err != nil {
		http.Error(w, "Failed to decode sessions from DB", http.StatusInternalServerError)
		log.Printf("Error decoding sessions: %v", err)
		return
	}

	if sessions == nil {
		sessions = []CheckSession{}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(sessions)
}

func getSessionByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid Session ID format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session CheckSession
	err = sessionCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&session)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "Session not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(session)
}

func deleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	objID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid Session ID format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := sessionCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Session deleted successfully",
		"id":      idStr,
	})
}
