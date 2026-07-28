package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type Ball struct {
	hits       int
	lastPlayer Player
}

type Player struct {
	name string
	key  string
}

func player(player Player, teamChan chan Ball, oppositeChan chan Ball, stopChan chan Ball, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stopChan:
			return

		case ball, ok := <-teamChan:
			if !ok {
				return
			}

			if ball.lastPlayer == player {
				select {
				case teamChan <- ball:
				case <-stopChan:
					return
				}
				time.Sleep(1 * time.Second)
				continue
			}

			if rand.Float32() < 0.02 {
				fmt.Printf("[%s] didn't catch the ball (%d hits)\n", player.name, ball.hits)
				close(stopChan)
				return
			}

			ball.hits++
			ball.lastPlayer = player
			fmt.Printf("[%s] hit the ball (%d hits)\n", player.name, ball.hits)

			time.Sleep(1 * time.Second)

			select {
			case oppositeChan <- ball:
			case <-stopChan:
				return
			}
		}
	}
}

func runSingleMatch() {
	fmt.Println("[single] start match")

	var wg sync.WaitGroup
	pingChan := make(chan Ball)
	pongChan := make(chan Ball)
	stopChan := make(chan Ball)

	wg.Add(2)
	go player(Player{"ping", ""}, pingChan, pongChan, stopChan, &wg)
	go player(Player{"pong", ""}, pongChan, pingChan, stopChan, &wg)

	fmt.Println("[single] first ball")
	pingChan <- Ball{0, Player{}}

	wg.Wait()
	fmt.Println("[single] match is over")
}

func runDoubleMatch() {
	fmt.Println("[double] start match")

	var wg sync.WaitGroup
	teamAChan := make(chan Ball)
	teamBChan := make(chan Ball)
	stopChan := make(chan Ball)

	wg.Add(4)
	go player(Player{"ping", "A"}, teamAChan, teamBChan, stopChan, &wg)
	go player(Player{"pong", "A"}, teamAChan, teamBChan, stopChan, &wg)
	go player(Player{"peng", "B"}, teamBChan, teamAChan, stopChan, &wg)
	go player(Player{"pang", "B"}, teamBChan, teamAChan, stopChan, &wg)

	fmt.Println("[double] first ball")
	teamAChan <- Ball{0, Player{}}

	wg.Wait()
	fmt.Println("[double] match is over")
}

//func runSimulation(input *bufio.Reader) { }

func mainMenu(input *bufio.Reader) {
	for {
		fmt.Println("\n[menu] game select menu")
		fmt.Println("1) single match (1v1)")
		fmt.Println("2) double match (2v2)")
		fmt.Println("3) simulate 'n' single matches")
		fmt.Print("[menu] enter command (type 'no' to cancel): ")

		command, _ := input.ReadString('\n')
		command = strings.TrimSpace(command)
		fmt.Print("\n")

		switch command {
		case "1":
			runSingleMatch()
		case "2":
			runDoubleMatch()
		case "3":
			// runSimulation()
		case "no":
			return
		default:
			fmt.Println("[menu] unknown command")
		}
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	mainMenu(input)
}
