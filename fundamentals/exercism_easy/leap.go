package main

import "fmt"

func main() {
	var year int
	_, err := fmt.Scan(&year)
	if err != nil {
		return
	}

	if year%4 == 0 || year%400 == 0 {
		fmt.Printf("Year %d is a leap year", year)
		return
	}

	fmt.Printf("Year %d is not a leap year", year)
	return
}
