// Quadratic Equation Calculator
// James Taddei
// 4/3/21

// Package Imports
package main

// Import Statements
import (
	"fmt"
	"math"
)

// Main function, prints the two possible roots
func main() {
	// Has the user input if there's fractions or not
	fmt.Print("Are there fractions (y/n)? ")
	var fractions string
	fmt.Scan(&fractions)

	// Variable Creation
	var a float64
	var b float64
	var c float64

	// Has the user input A, B, and C based on if there are fractions
	switch fractions {
	case "y": // Fractions exist
		// Fraction Variable Inputs
		fmt.Println("Enter the variable with one of the following input: '3 1'")
		a = fraction_input('a')
		b = fraction_input('b')
		c = fraction_input('c')
	case "n": // No fractions
		// Normal Variable Inputs
		fmt.Print("Enter 'a': ")
		fmt.Scan(&a)
		fmt.Print("Enter 'b': ")
		fmt.Scan(&b)
		fmt.Print("Enter 'c': ")
		fmt.Scan(&c)
	default: // Wrong input, restarts program
		fmt.Println("Error, please enter either 'y' or 'n'. Restarting")
		main()
		return
	}

	// Finds the different chunk values of the equation
	part1 := -1 * b
	part2 := math.Sqrt((math.Pow(b, 2) - (4 * a * c)))
	part3 := 2 * a

	// Solves for the 2 different answers
	quadPlus := (part1 + part2) / part3
	quadMinus := (part1 - part2) / part3

	// Final prints statements of the 2 roots
	fmt.Printf("x = %f", quadPlus)
	fmt.Printf("\nx = %f", quadMinus)
}

// Function that has the user input a fraction number and returns the value
func fraction_input(variable byte) float64 {
	fmt.Printf("Enter '%c': ", variable)
	var input1 float64
	var input2 float64
	fmt.Scan(&input1)
	fmt.Scan(&input2)

	return (input1 / input2)
}

/* Notes:
Equation: Vx = (-b +- sqrt(b^2 - 4ac)) / 2a
*/
