// Factorial Finder
// James Taddei
// 4/5/21

// Package Imports
package main

// Import Statements
import (
	"fmt"
)

// Main function - finds and prints the factorial of an inputted number
func main() {
	// User input for number this will find the factorial of
	fmt.Print("Factorial of: ")
	var factorialOf int
	fmt.Scan(&factorialOf)

	factorial := 1

	// Goes through each integer from 1-n to find the factorial of n
	for i := 1; i <= factorialOf; i++ {
		factorial *= i
	}

	// Final print statment of the factorial of n
	fmt.Printf("The factorial of %d is: %d", factorialOf, factorial)
}
