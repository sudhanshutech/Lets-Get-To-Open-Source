package main

import "fmt"

func main() {
	fmt.Println("This is if else example")

	myAge := 25
	var Result string

	// brackets can not go to next line in Go
	if myAge >= 18 {
		Result = "Adult"
	} else {
		Result = "Not an adult"
	}

	fmt.Println("Result is ", Result)

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if sometimes values comes from web request or database then we can assign it to a variable and use it in if condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// if err != nil {
	// 	fmt.Println(err)
	// }

}
