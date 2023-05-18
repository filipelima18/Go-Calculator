package main

import (
	"errors"
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

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	x := 10
	y := 2

	sum := add(x, y)
	difference := subtract(x, y)
	mult := multiply(x, y)
	div, err := divide(x, y)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", mult)
	fmt.Println("Quotient:", div)
}