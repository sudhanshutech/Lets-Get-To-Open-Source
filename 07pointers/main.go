package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println("The value of ptr is ", ptr)

	myNumber := 23
	var ptr *int = &myNumber                             // this is how we assign the address of myNumber to ptr
	fmt.Println("The value of myNumber is ", ptr)        // this will print the address of myNumber
	fmt.Println("The value of actual pointer is ", *ptr) // this will print the value of myNumber

	*ptr = *ptr + 1                        // this will increment the value of myNumber by 1
	fmt.Println("The value is ", myNumber) // this will print the value of myNumber
}
