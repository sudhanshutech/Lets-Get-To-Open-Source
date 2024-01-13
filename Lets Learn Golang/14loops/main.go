package main

import "fmt"

func main() {
	fmt.Println("This is loop example")

	months := [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct"}

	fmt.Println(months)

	for i := 0; i < len(months); i++ { // for loop: traditional
		fmt.Println(months[i])
	}

	for index := range months { // for range loop: index only
		fmt.Println(months[index])
	}

	for index, value := range months { // for range loop: index and value
		fmt.Println(index, value)
	}

	//goto example

	myVal := 20

	for myVal < 30 {

		if myVal == 25 {
			myVal += 1
			goto Next
		}
		fmt.Println(myVal)
		myVal++
	}

Next:
	fmt.Println("This is next")

}

// Description of Loops

// Loops are used to execute a set of statements repeatedly until a particular condition is satisfied. Go has only one looping construct, the for loop.
//  The basic for loop has three components separated by semicolons: the init statement: executed before the first iteration, the condition expression: evaluated before every iteration and the post statement: executed at the end of every iteration.
