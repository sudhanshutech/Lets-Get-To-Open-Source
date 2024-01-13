package main

import "fmt"

func main() {

	fmt.Println("This is first function")
	func2()
	fmt.Println(AddTwo(10, 20))
	fmt.Println(AddAll(10, 20, 30, 40, 50, 1))
}

func func2() {

	fmt.Println("This is second function")
}

func AddTwo(val1 int, val2 int) int {
	return val1 + val2
}

func AddAll(values ...int) int { // variadic function: a function that can take variable number of arguments
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// Description of Functions

// Functions are a set of statements that are grouped together to perform a specific task. Functions help in code reusability.
