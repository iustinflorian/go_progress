package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	isValid := true

	fmt.Print("1st number: ")
	fmt.Scan(&num1)

	fmt.Print("Operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("2st number: ")
	fmt.Scan(&num2)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero not allowed")
			isValid = false
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator.")
		isValid = false
	}

	if isValid {
		fmt.Printf("Result: %.2f\n", result)
	}
}
