package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, gameChan chan Ball) {
	for ball := range gameChan {
		if rand.Float32() < 0.2 {
			fmt.Printf("[%s] didn't catch the ball after %d hits\n", name, ball.hits)
			close(gameChan)
			return
		}
		ball.hits++
		fmt.Printf("[%s] successfully hit the ball\n", name)
		time.Sleep(1 * time.Second)
		gameChan <- ball
	}
}

func main() {
	fmt.Println("[main] start match")

	gameChan := make(chan Ball)

	go player("ping", gameChan)
	go player("pong", gameChan)

	fmt.Println("[main] first ball")
	gameChan <- Ball{0}

	time.Sleep(8 * time.Second)
	fmt.Println("[main] match is over")
}
