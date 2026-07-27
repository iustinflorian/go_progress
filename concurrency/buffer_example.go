package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		fmt.Println("[goroutine] sending part1")
		ch <- "part1"
		fmt.Println("[goroutine] part1 in buffer")

		fmt.Println("[goroutine] sending part2...")
		ch <- "part2"
		fmt.Println("[goroutine] part2 in buffer")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("\n[main] start reading")
	fmt.Println("[main] done reading:", <-ch)

	time.Sleep(1 * time.Second)
	fmt.Println("[Main] done reading:", <-ch)
}