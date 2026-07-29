package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

func player(player Player, teamChan chan Ball, oppositeChan chan Ball, stopChan chan Ball, wg *sync.WaitGroup, show bool) {
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

			if rand.Float32() < 0.4 {
				if show {
					fmt.Printf("[%s] didn't catch the ball (%d hits)\n", player.name, ball.hits)
				}
				close(stopChan)
				return
			}

			ball.hits++
			ball.lastPlayer = player
			if show {
				fmt.Printf("[%s] hit the ball (%d hits)\n", player.name, ball.hits)
			}

			time.Sleep(1 * time.Second)

			select {
			case oppositeChan <- ball:
			case <-stopChan:
				return
			}
		}
	}
}

func runSingleMatch(show bool) {
	if show {
		fmt.Println("[single] start match")
	}

	var wg sync.WaitGroup
	pingChan := make(chan Ball)
	pongChan := make(chan Ball)
	stopChan := make(chan Ball)

	wg.Add(2)
	go player(Player{"ping", ""}, pingChan, pongChan, stopChan, &wg, show)
	go player(Player{"pong", ""}, pongChan, pingChan, stopChan, &wg, show)

	if show {
		fmt.Println("[single] first ball")
	}
	pingChan <- Ball{0, Player{}}

	wg.Wait()
	if show {
		fmt.Println("[single] match is over")
	}
}

func runDoubleMatch() {
	fmt.Println("[double] start match")

	var wg sync.WaitGroup
	teamAChan := make(chan Ball)
	teamBChan := make(chan Ball)
	stopChan := make(chan Ball)

	wg.Add(4)
	go player(Player{"ping", "A"}, teamAChan, teamBChan, stopChan, &wg, true)
	go player(Player{"pong", "A"}, teamAChan, teamBChan, stopChan, &wg, true)
	go player(Player{"peng", "B"}, teamBChan, teamAChan, stopChan, &wg, true)
	go player(Player{"pang", "B"}, teamBChan, teamAChan, stopChan, &wg, true)

	fmt.Println("[double] first ball")
	teamAChan <- Ball{0, Player{}}

	wg.Wait()
	fmt.Println("[double] match is over")
}

func runSimulation(input *bufio.Reader) {
	fmt.Print("[simulation] how many matches to simulate? ")

	command, _ := input.ReadString('\n')
	command = strings.TrimSpace(command)
	count, err := strconv.Atoi(command)
	if err != nil || count <= 0 {
		fmt.Println("[simulation] invalid match count")
		return
	}

	fmt.Printf("[simulation] starting %d matches\n", count)
	startTime := time.Now()

	var wg sync.WaitGroup

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runSingleMatch(false)
		}()
	}

	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("[simulation] finished %d matches in %s\n", count, duration)
}

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
			runSingleMatch(true)
		case "2":
			runDoubleMatch()
		case "3":
			runSimulation(input)
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
