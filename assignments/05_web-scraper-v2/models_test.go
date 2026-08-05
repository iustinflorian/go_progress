package main

import (
	"testing"
	"time"
)

func TestCheckSessionInitialization(t *testing.T) {
	results := []URLResult{
		{URL: "https://example.com", StatusCode: 200, Title: "Example"},
		{URL: "https://google.com", StatusCode: 200, Title: "Google"},
	}

	session := CheckSession{
		CreatedAt: time.Now(),
		TotalURLs: len(results),
		Results:   results,
	}

	if session.TotalURLs != 2 {
		t.Errorf("CheckSession TotalURLs = %d, want 2", session.TotalURLs)
	}

	if len(session.Results) != 2 {
		t.Errorf("CheckSession Results length = %d, want 2", len(session.Results))
	}
}
