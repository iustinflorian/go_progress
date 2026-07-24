package main

import (
	"fmt"
	"net/http"
	"os"
)

var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		location := os.Getenv("LOCATION")
		if location == "" {
			location = "Go"
		}

		count++

		fmt.Fprintf(w, "<h1>Hello from %s! This page has been visited %d times</h1>", location, count)
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
