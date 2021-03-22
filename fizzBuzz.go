// Go FizzBuzz
// James Taddei
// 3/15/21

package main

import (
	"fmt"
)

func main() {
	// Storing "Fizz" Number
	fmt.Print("Enter 'Fizz' number: ")
	var fizzNum int
	fmt.Scan(&fizzNum)

	// Storing "Buzz" Number
	fmt.Print("Enter 'Buzz' numer: ")
	var buzzNum int
	fmt.Scan(&buzzNum)

	// Storing Loop to Number
	fmt.Print("Enter loop to number: ")
	var maxNum int
	fmt.Scan(&maxNum)

	// Main Loop for Each Number in 0-Max
	for i := 0; i <= maxNum; i++ {
		if i%(fizzNum*buzzNum) == 0 {
			fmt.Println("FizzBuzz")
		} else if i%fizzNum == 0 {
			fmt.Println("Fizz")
		} else if i%buzzNum == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
