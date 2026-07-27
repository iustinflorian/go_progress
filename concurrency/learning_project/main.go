package main

import (
	"fmt"
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

func assemblePart(part Part, resultChan chan<- AssemblyResult) {
	//defer wg.Done()
	fmt.Printf("[goroutine] start assembly on %s\n", part.name)
	time.Sleep(part.duration)
	//fmt.Printf("[goroutine] %s assembly complete\n", part.name)

	resultChan <- AssemblyResult{
		partName: part.name,
		success:  true,
	}
}

func main() {
	fmt.Println("[main] start")

	//var wg sync.WaitGroup
	resultChan := make(chan AssemblyResult)

	parts := []Part{
		{"engine", 3 * time.Second},
		{"suspension", 3 * time.Second},
		{"wheel", 3 * time.Second},
	}

	for _, part := range parts {
		//wg.Add(1)
		go assemblePart(part, resultChan)
	}

	// go assembleEngine("engine", 5*time.Second)
	// go assembleEngine("suspension", 2*time.Second)
	// go assembleEngine("wheel", 3*time.Second)

	fmt.Println("[main] waiting for channel results")

	//wg.Wait()

	for i := 0; i < len(parts); i++ {
		result := <-resultChan
		fmt.Printf("[channel] %s assembly complete\n", result.partName)
	}

	fmt.Println("[main] end")
}
