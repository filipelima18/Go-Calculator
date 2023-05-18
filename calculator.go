package main

import "fmt"

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
	x := 40
	y := 4

	sum := add(x, y)
	difference := subtract(x, y)
	mult := multiply(x, y)
	div := divide(x, y)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", mult)
	fmt.Println("Quotient:", div)
}
