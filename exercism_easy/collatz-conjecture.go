package main

import "fmt"

func main() {
	var col func(n int) int
	steps := -1

	col = func(n int) int {
		steps++
		if n == 1 {
			return 1
		}
		if n%2 == 0 {
			return col(n / 2)
		}
		return col(n*3 + 1)
	}

	col(12)
	fmt.Println(steps)
}
