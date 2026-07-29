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
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	response, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		return Result{
			URL:      url,
			Err:      fmt.Errorf("-> timeout: %w", err),
			Duration: duration,
		}
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return Result{
			URL:        url,
			StatusCode: response.StatusCode,
			Err:        fmt.Errorf("-> http error %d: %s", response.StatusCode, http.StatusText(response.StatusCode)),
			Duration:   duration,
		}
	}

	return Result{
		URL:        url,
		StatusCode: response.StatusCode,
		Err:        nil,
		Duration:   duration,
	}
}

func worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("[worker-%d] is fetching %s\n", id, url)
		results <- fetch(url)
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtuble.com",
		"https://httpbin.org/status/404",
		"https://www.reddit.com",
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(urls))
	jobs := make(chan string, len(urls))

	numWorkers := 2
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w+1, jobs, results, &wg)
	}

	fmt.Printf("[main] fetching %d urls\n", len(urls))
	start := time.Now()

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

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

	fetchTime := time.Since(start)
	fmt.Printf("[main] it took %s for %d workers to process %d urls\n", fetchTime, numWorkers, len(urls))
}
