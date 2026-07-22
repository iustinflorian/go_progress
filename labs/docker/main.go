package main

import (
	"fmt"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Fprintf(w, "<h1>Hello from Go! This page has been visited %d times</h1>", count)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
