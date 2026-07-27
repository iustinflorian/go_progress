package main

import (
	"fmt"
	"unicode"
)

/*
Count the frequency of letters in texts using parallel computation.
Parallelism is about doing things in parallel that can also be done sequentially.
A common example is counting the frequency of letters.
Employ parallelism to calculate the total frequency of each letter in a list of texts.
*/

func Frequency(s string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			freqMap[letter]++
		}
	}
	return freqMap
}

func main() {
	text := []string{
		"this is the first string",
		"this is the second string",
		"this is the third string",
	}

	var finalMap = make(map[rune]int)
	mapChan := make(chan map[rune]int)

	for _, str := range text {
		go func(s string) {
			mapChan <- Frequency(s)
		}(str)
	}

	for i := 0; i < len(text); i++ {
		partialMap := <-mapChan
		for letter, count := range partialMap {
			finalMap[letter] += count
		}
	}

	fmt.Printf("Frequency of letter '%s' is: %d\n", "e", finalMap['e'])
	fmt.Printf("Frequency of letter '%s' is: %d\n", "a", finalMap['a'])
	fmt.Printf("Frequency of letter '%s' is: %d\n", "t", finalMap['t'])
}
