package main

import (
	"fmt"
)

// Function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int

	// Input from user
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Calculate factorial
	result := factorial(num)

	// Display the result
	fmt.Printf("Factorial of %d is %d\n", num, result)
}
