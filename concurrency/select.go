package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	strings := make(chan string)
	nums := make(chan int)

	go func() {
		for {
			strings <- "Hello"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			nums <- 1
			time.Sleep(time.Second * 2)
		}
	}()

	/*
		for {
			fmt.Println(<-strings)
			// This will block waiting
			fmt.Println(<-nums)
		}
	*/

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	for {
		select {
		case msg := <-strings:
			fmt.Println(msg)
		case msg := <-nums:
			fmt.Println(msg)
		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
			return
		}
	}
}
