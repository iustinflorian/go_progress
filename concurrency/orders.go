package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func processOrder(waiterID int, order Order, ordChan chan<- string) string {
	fmt.Printf("Waiter %d: Preparing order for table %d...\n", waiterID, order.TableNumber)
	time.Sleep(order.PrepTime)
	fmt.Printf("Waiter %d: Order ready for table %d!\n\n", waiterID, order.TableNumber)

	ord := fmt.Sprintf("Order for table %d is done!\n", order.TableNumber)
	ordChan <- ord

	return "Order is done!"
}

func main() {
	orders := []Order{
		{1, 2 * time.Second},
		{2, 3 * time.Second},
		{3, 4 * time.Second},
		{4, 2 * time.Second},
		{5, 1 * time.Second},
	}

	wg := sync.WaitGroup{}
	ordChan := make(chan string)
	var responses []string

	for waiterID, order := range orders {
		wg.Add(1)

		go func() {
			defer wg.Done()
			processOrder(waiterID, order, ordChan)
			//responses = append(responses, processOrder(waiterID, order))
			//wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ordChan)
	}()

	for ord := range ordChan {
		responses = append(responses, ord)
	}

	fmt.Println(responses)
	//fmt.Println("All orders ready!")
}
