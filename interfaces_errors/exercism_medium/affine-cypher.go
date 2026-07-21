package main

import (
	"fmt"
	"strings"
)

// encryption func: e = (a * chIndex + b) % m
// a, m coprime (or, relatively prime)
// decryption func: d = (modInverse(a, m) * ((chIndex - b) % m + m)) % m

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

func modInverse(a, m int) int {
	a = a % m
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return 1
}

func Encrypt(text string, a, b, m int) (string, error) {
	var result []rune

	if a < 1 {
		return "", ErrInvalidNum(a)
	}
	if !checkCoprime(a, m) {
		return "", ErrNotCoprime(a)
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

	return string(result), nil
}

func Decrypt(encryptedText string, a, b, m int) (string, error) {
	var result []rune
	encryptedText = strings.ToLower(encryptedText)

	if a < 1 {
		return "", ErrInvalidNum(a)
	}
	if !checkCoprime(a, m) {
		return "", ErrNotCoprime(a)
	}

	for _, ch := range encryptedText {
		if ch >= 'a' && ch <= 'z' {
			chIndex := int(ch - 'a')
			deIndex := (modInverse(a, m) * ((chIndex-b)%m + m)) % m
			deChar := rune('a' + deIndex)

			result = append(result, deChar)
		}
	}

	return string(result), nil
}

func main() {
	a, b, m := 5, 8, 26
	text := "hello there"

	encrypted, err := Encrypt(text, a, b, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encrypted)

	decrypted, err := Decrypt(encrypted, a, b, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decrypted)
}
