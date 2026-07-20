package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input word: ")
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	words := strings.Fields(s)

	var acronym string

	for _, word := range words {
		letter := string(word[0])
		acronym += letter
	}

	fmt.Println(strings.ToUpper(acronym))
}
