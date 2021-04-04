// Gravity Calculator
// James Taddei
// 4/3/21

// Main Package Import
package main

// Module Imports
import (
	"fmt"
	"math"
)

// Main function, finds the the object's height after 10 seconds
func main() {
	// Variable Creation
	gravity := -9.81                // a in formula, Earth's gravity in m/s^2
	var initialVelocity float64 = 0 // vi in formula, should be 0
	var fallingTime float64 = 10    // t in formula, should be 10 (s)
	var initialPosition float64 = 0 // xi in formula, should be 0
	var finalPosition float64       // x(t) in formula, final position the program solves for

	// Calculation + Final Print Statement
	finalPosition = 0.5*(gravity*math.Pow(fallingTime, 2)) + (initialVelocity * fallingTime) + initialPosition
	fmt.Printf("The object's position after %f seconds is %f m.", fallingTime, finalPosition) // Prints the objects final height
}

/* Notes:
Formula: x(t) = 0.5 * at^2 + vi*t + xi
Note: The correct value is -490.5 m. Golang will output more digits after the decimal place, but that is unimportant
Final Print: height after 10 secs
*/
