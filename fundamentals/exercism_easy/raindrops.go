package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var res string
	notDivisible := true

	fmt.Print("Input n: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	if n%3 == 0 {
		res = res + "Pling"
		notDivisible = false
	}
	if n%5 == 0 {
		res = res + "Plang"
		notDivisible = false
	}
	if n%7 == 0 {
		res = res + "Plong"
		notDivisible = false
	}
	if notDivisible {
		res = strconv.Itoa(n)
	}

	fmt.Println("Result: ", res)
}
