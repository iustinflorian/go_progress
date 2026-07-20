package main

import "fmt"

func main() {
	n := 10
	sum := 0
	squareSum := 0

	for i := 1; i <= n; i++ {
		sum += i
		squareSum += i * i
	}

	fmt.Println(sum*sum - squareSum)
}
