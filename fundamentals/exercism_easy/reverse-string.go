package main

import "fmt"

func main() {
	var s string
	fmt.Print("Input s: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println(err)
		return
	}

	reverse := ""

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	fmt.Println(reverse)
}
