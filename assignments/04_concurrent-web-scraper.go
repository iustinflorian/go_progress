package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Err        error
	Duration   time.Duration
}

func fetch(url string) Result {
	fmt.Printf("[goroutine] fetching %s\n", url)

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
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.reddit.com",
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(urls))

	fmt.Printf("[main] fetching %d urls\n", len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func() {
			defer wg.Done()

			results <- fetch(url)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Err != nil {
			fmt.Printf("[error] fetching %s ended with error:\n%v (%v)\n", res.URL, res.Err, res.Duration)
		} else {
			fmt.Printf("[ok %d] fetched %s (%v)\n", res.StatusCode, res.URL, res.Duration)
		}
	}
}
