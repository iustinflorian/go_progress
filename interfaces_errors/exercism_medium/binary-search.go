package main

import "fmt"

func binarySearch(list []int, value int) int {
	i := 0
	j := len(list) - 1
	for i < j {
		middleIndex := i + (j-i)/2
		middle := list[middleIndex]
		if value == middle {
			return middleIndex
		} else if value > middle {
			i = middle + 1
		} else if value < middle {
			j = middle - 1
		}
	}
	return -1
}

func main() {
	list := []int{2, 4, 12, 26, 34, 56, 79, 90, 111}
	value := 34

	if binarySearch(list, value) < 0 {
		fmt.Println("Not found")
		return
	}
	fmt.Printf("The list contains number %d at index %b\n", value, binarySearch(list, value))
}
