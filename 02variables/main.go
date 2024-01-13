package main

import "fmt"

func main() {
	var username string = "my name" // string is a sequence of characters
	fmt.Println("Hello, ", username, "!")
	fmt.Printf("Variable type: %T\n", username)

	var age uint8 = 30 // uint8 is an 8-bit unsigned integer
	fmt.Println("My age is", age)
	fmt.Printf("Variable type: %T\n", age)

	var temperature float32 = 25.6333439433 // float32 is a 32-bit floating point number
	fmt.Println("The temperature is", temperature)
	fmt.Printf("Variable type: %T\n", temperature)

	var myvariable int // default value is 0
	fmt.Println("My variable is", myvariable)
	fmt.Printf("Variable type: %T\n", myvariable)

	//implicit type
	var myvariable2 = 10
	fmt.Println("My variable is", myvariable2)
	fmt.Printf("Variable type: %T\n", myvariable2)

	//no var keyword
	myvariable3 := 10
	fmt.Println("My variable is", myvariable3)
}
