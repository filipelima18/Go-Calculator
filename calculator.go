package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func main() {
	var operation string
	var x, y int

	fmt.Print("Enter the operation (+, -, *, /): ")
	fmt.Scan(&operation)
	fmt.Print("Enter the first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&y)

	var result int

	switch operation {
	case "+":
		result = add(x, y)
	case "-":
		result = subtract(x, y)
	case "*":
		result = multiply(x, y)
	case "/":
		result = divide(x, y)
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println("Result:", result)
}
