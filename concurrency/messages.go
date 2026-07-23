package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessage(num int, msgChan chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Sending message %d\n", num)
	time.Sleep(time.Duration(num) * time.Second)
	msgChan <- fmt.Sprintf("Message %d sent!", num)
	wg.Done()
}

func receiveMessage(msgChan <-chan string) {
	fmt.Printf("Receiving message\n")

	for msg := range msgChan {
		fmt.Printf("Received message %s\n", msg)
	}
}

func main() {
	msgChan := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go sendMessage(3, msgChan, &wg)
	go sendMessage(5, msgChan, &wg)

	go func() {
		receiveMessage(msgChan)
	}()

	wg.Wait()
	close(msgChan)

	fmt.Println("Done")
	fmt.Scanln()
}
