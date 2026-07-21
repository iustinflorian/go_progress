package main

import (
	"fmt"
	"strings"
)

// encryption func: e = (a * i + b) % m
// a, m coprime (or, relatively prime)

type ErrInvalidNum int
type ErrNotCoprime int

func (e ErrInvalidNum) Error() string {
	return fmt.Sprintf("Value of 'a' (in this case %v) should be greater than 0.", int(e))
}
func (e ErrNotCoprime) Error() string {
	return fmt.Sprintf("Value of 'a' (in this case %v) and 'm' not coprime.", int(e))
}

func checkCoprime(a, m int) bool {
	for i := 2; i <= a; i++ {
		if a%i == 0 && m%i == 0 {
			return false
		}
	}
	return true
}

func Encrypt(text string, a, b, m int) ([]rune, error) {
	var result []rune

	if a < 1 {
		return nil, ErrInvalidNum(a)
	}
	if !checkCoprime(a, m) {
		return nil, ErrNotCoprime(a)
	}

	text = strings.ToLower(text)

	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			chIndex := int(ch - 'a')
			enIndex := (a*chIndex + b) % m
			enChar := rune('a' + enIndex)

			result = append(result, enChar)
		}
	}

	return result, nil
}

func main() {
	a, b, m := 5, 8, 26
	text := "hello"

	encrypted, err := Encrypt(text, a, b, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encrypted))
}
