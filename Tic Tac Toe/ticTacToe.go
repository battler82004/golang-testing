// Tic Tac Toe
// James Taddei
// 4/3/21

// Package Imports
package main

// Import Statements
import (
	"fmt"
	"math/rand"
	"time"
)

// Main function which decides if bot or player goes first, who wins, and prints out the winner
func main() {
	// Beginning prints and player or bot first input
	fmt.Println("This is a tic-tac-toe game created by James Taddei in golang.")
	fmt.Print("Type '1' to go first and '2' to go second: ")
	var playerPos int
	fmt.Scan(&playerPos)

	// Branches based on if the player or bot goes first, the called function starts the actual game
	var onBoard [9]string
	switch playerPos {
	case 1: // Goes through the game with the player having the first turn
		onBoard = player_first()
	case 2: // Goes through the game with the bot having the first turn
		onBoard = bot_first()
	default: // Restarts the program if the wrong number is input
		fmt.Println("Error, please enter either '1' or '2'. Restarting program.")
		main()
	}

	// Finds the number of X's and O's to help find who wins
	xCount := 0
	oCount := 0
	for i := 0; i < 9; i++ {
		if onBoard[i] == "X" {
			xCount++
		} else if onBoard[i] == "O" {
			oCount++
		}
	}

	// Determines who wins based on who went first and number of X's and O's
	var winner string
	if win_checker(onBoard) {
		switch playerPos {
		case 1: // Player First
			if xCount > oCount {
				winner = "player"
			} else {
				winner = "bot"
			}
		case 2: // Bot First
			if oCount > xCount {
				winner = "bot"
			} else {
				winner = "player"
			}
		}
	} else { // Winner statement for a draw (since winchecker <- false)
		winner = "none"
	}

	// Final print statement based on who won
	switch winner {
	case "player":
		fmt.Println("\n\nYou have triumphed over the bot! Congratulations on the victory!")
	case "bot":
		fmt.Println("\n\nCongratulations on your defeat against a bot who's not even programmed always even stop your win!")
	case "none":
		fmt.Println("\n\nYou have tied with the bot. Kind of anticlimactic. I don't know if I'm supposed to congratulate you for not losing or make fun of you for not winning.")
	}
}

// Prints a board with numbers on it so the player understands the numbering system
func turn1() {
	fmt.Println("\nYou will place 'X's based on the number position below:")
	num1To9 := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	board_printer(num1To9)
}

// Loop of all turns if the player goes first
func player_first() [9]string {
	// Creates 'onBoard' (where the board is stored) and prints a board with numbers on it
	onBoard := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	turn1()

	// Variable Creation for Loop
	won := false
	var playerTurnLocation int
	var count int = 0
	// Loop for each actual turn of game
	for !won {
		// Has the player input their move location
		fmt.Print("Location to place 'X': ")
		fmt.Scan(&playerTurnLocation)
		playerTurnLocation--

		// Checks if the player's input is valid
		if valid_location(onBoard, playerTurnLocation) != "no_error" {
			if playerTurnLocation > 8 || playerTurnLocation < 0 {
				fmt.Println("Error, number should be from 1-9")
			} else if onBoard[playerTurnLocation] != " " {
				fmt.Println("Error, something is already there")
			}
			continue
		}

		// Actually adds the player's move
		onBoard[playerTurnLocation] = "X"
		board_printer(onBoard)
		won = win_checker(onBoard)
		if won {
			break
		}

		// Checks for a draw by seeing if each position is filled
		count = 0
		for i := 0; i < 9; i++ {
			if onBoard[i] != " " {
				count++
			}
		}
		if count == 9 {
			break
		}

		// Has the bot go through its own move
		onBoard = bot_turn(onBoard)
		board_printer(onBoard)
		won = win_checker(onBoard)
	}

	// Final return statement of the final board
	return onBoard
}

// Loop of all turns if the bot goes first
func bot_first() [9]string {
	// Creates 'onBoard' (where the board is stored) and prints a board with numbers on it
	onBoard := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
	turn1()

	// Variable Creation for Loop
	won := false
	var playerTurnLocation int
	var count int = 0
	playerTurn := false
	// Loop for each actual turn of game
	for !won {
		// If the player didn't input an invalid # before, has the bot go through its own move
		if !playerTurn {
			onBoard = bot_turn(onBoard)
			board_printer(onBoard)
			won = win_checker(onBoard)
			if won {
				break
			}
		}

		// Checks for a draw by seeing if each position is filled
		count = 0
		for i := 0; i < 9; i++ {
			if onBoard[i] != " " {
				count++
			}
		}
		if count == 9 {
			break
		}

		// Has the player input their move location
		fmt.Print("Location to place 'X': ")
		fmt.Scan(&playerTurnLocation)
		playerTurnLocation--

		// Checks if the player's input is valid
		if valid_location(onBoard, playerTurnLocation) != "no_error" {
			if playerTurnLocation > 8 || playerTurnLocation < 0 {
				fmt.Println("Error, number should be from 1-9")
			} else if onBoard[playerTurnLocation] != " " {
				fmt.Println("Error, something is already there")
			}
			playerTurn = true
			continue
		}

		// Actually adds the player's move
		playerTurn = false
		onBoard[playerTurnLocation] = "X"
		board_printer(onBoard)
		won = win_checker(onBoard)
	}

	// Final return statement of the final board
	return onBoard
}

// Prints out the current board filled with X's, O's, and spaces
func board_printer(onBoard [9]string) {
	fmt.Println("\n ", onBoard[0], "|", onBoard[1], "|", onBoard[2])
	fmt.Println("-------------")
	fmt.Println(" ", onBoard[3], "|", onBoard[4], "|", onBoard[5])
	fmt.Println("-------------")
	fmt.Println(" ", onBoard[6], "|", onBoard[7], "|", onBoard[8])
}

// Checks if the input location is a valid move
func valid_location(onBoard [9]string, location int) string {
	// Checks if the number would actually be on the board
	if location > 8 || location < 0 {
		return "error"
	} else if onBoard[location] != " " { // Checks if the location isn't already filled
		return "error"
	} else {
		return "no_error" // Final return statement for if the number is valid
	}
}

// Checks if someone has won and returns a boolean
func win_checker(onBoard [9]string) bool {
	// This var is all of the locations that need to be checked to look over the 8 possible ways to win
	// 8 checks - 3 vertical, 3 horizontal, 2 crosses
	locationsToCheck := [8][3]int{{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 4, 8}, {2, 4, 6}}

	// Variable Creation (each pos is just a way to write which part of 'locationsToCheck' we're using)
	var pos1 int
	var pos2 int
	var pos3 int

	// Goes through each possible way to win and checks if someone has
	for i := 0; i < 8; i++ {
		pos1 = locationsToCheck[i][0]
		pos2 = locationsToCheck[i][1]
		pos3 = locationsToCheck[i][2]
		if onBoard[pos1] == onBoard[pos2] && onBoard[pos2] == onBoard[pos3] && onBoard[pos1] != " " {
			return true
		}
	}

	// Final return statement which says that no one has won
	return false
}

// Has the bot go through a turn and return the new board
func bot_turn(onBoard [9]string) [9]string {
	// Variable Creation
	placeIn := can_win(onBoard, "X") // All of the places the bot could place an "O" in to stop the player from winning
	rand.Seed(time.Now().UnixNano()) // Generates a seed for random functions based on the current time
	var location int

	// Finds the location the bot will place its "O"
	if len(placeIn) >= 1 { // This section activates if the player can win and has a 67% chance of protecting the grid square
		if rand.Intn(3) == 0 {
			location = random_location(onBoard)
		} else { // Has the bot protect the grid square where the player could win
			location = placeIn[0]
		}
	} else { // Has the bot place the "O" in a random (valid) position
		location = random_location(onBoard)
	}

	// Actually places the "O"
	onBoard[location] = "O"

	// Final return statement of the new board
	return onBoard
}

// Generates a "random" valid location for the bot to use
func random_location(onBoard [9]string) int {
	// Sets the seed based on the current time
	rand.Seed(time.Now().UnixNano())

	// Finds a random valid number
	location := rand.Intn(9)
	for valid_location(onBoard, location) == "error" {
		location = rand.Intn(9)
	}

	// Final return statement of a random valid number
	return location
}

// Finds and returns any position(s) where the bot could place its "O" to stop the player from winning
func can_win(onBoard [9]string, checkingFor string) []int {
	// This var is all of the locations that need to be checked to look over the 8 possible ways to win
	// 8 checks - 3 vertical, 3 horizontal, 2 crosses
	locationsToCheck := [8][3]int{{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 4, 8}, {2, 4, 6}}
	var posToBlockWin []int

	// Goes through each possible way to win and checks if someone could
	for i := 0; i < 8; i++ {
		posToBlockWin = append(posToBlockWin, can_win_subfunction(onBoard, checkingFor, locationsToCheck[i]))
	}

	// Removes 'i' (any 10's, which were used as place holders) from 'posToBlockWin'
	for i := 0; i < len(posToBlockWin); i++ {
		if posToBlockWin[i] == 10 {
			posToBlockWin = append(posToBlockWin[:i], posToBlockWin[(i+1):]...)
			i--
		}
	}

	// Returns the positions where the bot should place its "O" to stop the player from winning
	return posToBlockWin
}

// Finds the position where the bot should place an "O" to stop the player from a specific win
func can_win_subfunction(onBoard [9]string, checkingFor string, posChecking [3]int) int {
	if onBoard[posChecking[0]] == onBoard[posChecking[1]] && onBoard[posChecking[0]] == checkingFor && onBoard[posChecking[2]] == " " {
		return posChecking[2] // Returns the 3rd element if that position is where the player could win
	} else if onBoard[posChecking[1]] == onBoard[posChecking[2]] && onBoard[posChecking[1]] == checkingFor && onBoard[posChecking[0]] == " " {
		return posChecking[0] // Returns the 1st element if that position is where the player could win
	} else if onBoard[posChecking[0]] == onBoard[posChecking[2]] && onBoard[posChecking[0]] == checkingFor && onBoard[posChecking[1]] == " " {
		return posChecking[1] // Returns the 2nd element if that position is where the player could win
	} else {
		return 10 // Returns 10 if there is no position where the player could win (for this specific way to win)
	}
}
