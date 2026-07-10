package main

import "fmt"

func checkIsogram(freq map[string]int, s string) {
	for i := 0; i < len(s); i++ {
		if freq[string(s[i])] == 1 {
			fmt.Printf("The word %s is NOT an isogram.\n", s)
			return
		}
		freq[string(s[i])]++
	}
	fmt.Printf("The word %s is an isogram.\n", s)
	return
}

func main() {
	var freq = map[string]int{}
	var s string

	fmt.Print("Input s: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}

	checkIsogram(freq, s)
}
