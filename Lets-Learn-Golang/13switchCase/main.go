package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed((time.Now().UnixNano())) // seed means the starting point of random number the random number generator, and unixNano returns the current time in nanoseconds
	diceNumber := rand.Intn(6) + 1     // rand.Intn(6) returns a random number between 0 and 5, so we add 1 to get a number between 1 and 6
	fmt.Println("You rolled a ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 space")
	case 2:
		fmt.Println("You can move 2 spaces")
	case 3:
		fmt.Println("You can move 3 spaces")
	case 4:
		fmt.Println("You can move 4 spaces")
	case 5:
		fmt.Println("You can move 5 spaces")
	case 6:
		fmt.Println("You can move 6 spaces")
		fallthrough // fallthrough means that the next case will be executed even if it doesn't match the switch condition
	default:
		fmt.Println("Something went wrong")
	}
}

// Description of Switch Case

// Switch case is used to execute a block of code based on the value of a variable.
//  It is similar to the switch statement in other programming languages.
// The switch statement in Go is used to evaluate a variable or an expression and execute the corresponding case. The switch statement can be used to replace multiple if else statements.
//  The switch statement in Go is different from other programming languages. In Go, the break statement is not required after each case. The break statement is automatically added after each case.
//  If we want to execute the next case even if it doesn't match the switch condition, we can use the fallthrough statement. The fallthrough statement is used to execute the next case even if it doesn't match the switch condition. The fallthrough statement should be the last statement in a case.
