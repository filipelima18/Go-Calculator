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

func main() {
	x := 10
	y := 2

	sum := add(x, y)
	difference := subtract(x, y)
	mult := multiply(x, y)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", mult)
}
