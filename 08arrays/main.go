package main

import "fmt"

func main() {

	var myArray [5]int
	var myArray2 [5]string

	myArray[0] = 1
	myArray[1] = 2
	myArray[3] = 4

	myArray2[0] = "apple"
	myArray2[1] = "banana"
	myArray2[4] = "mango"

	fmt.Println(myArray)
	fmt.Println(myArray2)

	fmt.Println("The length of myArray is ", len(myArray))
}
