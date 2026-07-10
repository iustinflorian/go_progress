package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	n = 153

	count := 0
	aux := n
	sum := 0

	for aux > 0 {
		count++
		aux /= 10
	}

	aux = n
	for aux > 0 {
		digit := aux % 10
		sum += int(math.Pow(float64(digit), float64(count)))
		aux /= 10
	}

	if sum == n {
		fmt.Println("Armstrong number")
	} else {
		fmt.Println("NOT Armstrong number")
	}
}
