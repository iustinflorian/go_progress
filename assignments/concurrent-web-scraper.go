package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Err        error
	Duration   time.Duration
}

func fetch(url string) Result {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	response, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		return Result{
			URL:      url,
			Err:      err,
			Duration: duration,
		}
	}
	defer response.Body.Close()

	return Result{
		URL:        url,
		StatusCode: response.StatusCode,
		Err:        err,
		Duration:   duration,
	}
}

func main() {
	res := fetch("https://www.conventionalcommits.org/en/v1.0.0/")

	if res.Err != nil {
		fmt.Printf("[error] fetching %s ended with error:\n%v (%v)\n", res.URL, res.Err, res.Duration)
	} else {
		fmt.Printf("[ok %d] fetched %s (%v)\n", res.StatusCode, res.URL, res.Duration)
	}
}
