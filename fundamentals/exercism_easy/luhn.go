package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func luhnFormula(s string) bool {
	sin := strings.TrimSpace(s)
	sin = strings.ReplaceAll(s, " ", "")

	sinRune := []rune(sin)
	for i := len(sinRune) - 1; i > 0; i = i - 2 {
		number := int(sinRune[i] - '0')
		if number*2 > 9 {
			number = number*2 - 9
			sinRune[i] = rune(number + '0')
		} else {
			number = number * 2
			sinRune[i] = rune(number + '0')
		}
	}
	sum := 0
	for i := range sinRune {
		number := int(sinRune[i] - '0')
		sum += number
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func main() {
	var s string

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	if luhnFormula(s) == true {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}
}
