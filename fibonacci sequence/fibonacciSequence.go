// Fibanacci Sequence
// James Taddei
// 4/4/21

// Package Imports
package main

// Import Statments
import (
	"fmt"
)

// Main funciton which prints the fibonacci sequence up to n (and the nth digit)
func main() {
	// Final number input
	fmt.Print("Enter the number you would like to go to: ")
	var n int
	fmt.Scan(&n)

	// Variable Creation
	var fibSequence []int
	var lastNum int = 0
	var currNum int = 1
	var nextNum int

	// Fibonacci Sequence Finder
	for len(fibSequence) < (n - 2) {
		nextNum = lastNum + currNum
		lastNum = currNum
		currNum = nextNum
		fibSequence = append(fibSequence, nextNum)
	}

	// Final prints statments based on n
	if n == 0 { // Print for 0 digits
		fmt.Println("\nThere is no sequence with 0 digits.")
	} else if n == 1 { // Print for 1 digit
		fmt.Printf("\nThe number in the 1st place is: 0")
	} else if n == 2 { // Print for 2 digits
		fmt.Println("\nThe sequence is: 0 and 1")
		fmt.Printf("The number in the 2nd place is: 1")
	} else { // Print for >2 digits
		var finalString string = "0, 1, "

		// Turns the slice of the sequence into a ", " separated string
		for i := 0; i < len(fibSequence)-1; i++ {
			finalString += fmt.Sprint(fibSequence[i], ", ")
		}

		// Prints the sequence and the nth number
		finalString += fmt.Sprint("and ", fibSequence[len(fibSequence)-1])
		fmt.Println("\nThe sequence is:", finalString)
		fmt.Printf("The number in place %d is: %d", n, fibSequence[len(fibSequence)-1])
	}
}
