package main

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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
