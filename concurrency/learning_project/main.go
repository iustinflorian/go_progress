package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Part struct {
	name     string
	duration time.Duration
}

type AssemblyResult struct {
	partName string
	success  bool
}

func assemblePart(wg *sync.WaitGroup, part Part, resultChan chan<- AssemblyResult, errorChan chan<- error) {
	defer wg.Done()
	fmt.Printf("[goroutine] start assembly on %s\n", part.name)
	time.Sleep(part.duration)

	if rand.Float32() < 0.4 {
		errorChan <- fmt.Errorf("assembly on %s failed", part.name)
		return
	}

	resultChan <- AssemblyResult{
		partName: part.name,
		success:  true,
	}
}

func main() {
	fmt.Println("[main] start")

	var wg sync.WaitGroup
	resultChan := make(chan AssemblyResult)
	errorChan := make(chan error)

	parts := []Part{
		{"engine", 3 * time.Second},
		{"suspension", 3 * time.Second},
		{"wheel", 3 * time.Second},
	}

	for _, part := range parts {
		wg.Add(1)
		go assemblePart(&wg, part, resultChan, errorChan)
	}

	fmt.Println("[main] waiting for channel results")

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
		fmt.Println("[main] done, channel closed")
	}()

	for {
		select {
		case res, ok := <-resultChan:
			if ok {
				fmt.Printf("[success] assembly on %s finished\n", res.partName)
			} else {
				resultChan = nil
			}
		case err, ok := <-errorChan:
			if ok {
				fmt.Printf("[error] %s\n", err)
			} else {
				errorChan = nil
			}
		}

		if resultChan == nil && errorChan == nil {
			break
		}
	}

	fmt.Println("[main] all results and errors processed")
	fmt.Println("[main] end")
}
