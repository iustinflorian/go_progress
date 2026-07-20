package main

import "fmt"

func main() {
	fmt.Print("Input name: ")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}

	if name == "" {
		fmt.Print("One for you, one for me")
	} else {
		fmt.Printf("One for %s, one for me\n", name)
	}
}
