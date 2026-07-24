package main

import (
	"fmt"
	"sync"
	"time"
)

type Part struct {
	name     string
	duration time.Duration
}

func assembleEngine(wg *sync.WaitGroup, part Part) {
	defer wg.Done()
	fmt.Printf("[goroutine] start assembly on %s\n", part.name)
	time.Sleep(part.duration)
	fmt.Printf("[goroutine] %s assembly complete\n", part.name)
}

func main() {
	fmt.Println("[main] start")

	var wg sync.WaitGroup

	parts := []Part{
		{"engine", 3 * time.Second},
		{"suspension", 3 * time.Second},
		{"wheel", 3 * time.Second},
	}

	for _, part := range parts {
		wg.Add(1)
		go assembleEngine(&wg, part)
	}

	// go assembleEngine("engine", 5*time.Second)
	// go assembleEngine("suspension", 2*time.Second)
	// go assembleEngine("wheel", 3*time.Second)

	fmt.Println("[main] still runs")

	wg.Wait()

	fmt.Println("[main] end")
}
