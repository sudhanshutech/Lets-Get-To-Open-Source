package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice = []int{1.0, 2.0, 3.0, 4.0, 5.0} // here we are using a slice which is a dynamic array
	fmt.Println(mySlice)
	fmt.Println("The length of mySlice is ", len(mySlice))

	mySlice = append(mySlice, 6.0, 7.0, 8.0, 9.0, 10.0) // we can append to a slice
	fmt.Println(mySlice)

	mySlice = append(mySlice[2:5]) // we can slice a slice
	fmt.Println(mySlice)

	newValues := make([]int, 5) // we can create a slice with make])

	newValues[0] = 10
	newValues[1] = 40
	newValues[2] = 30
	newValues[3] = 60
	newValues[4] = 50

	fmt.Println(newValues)
	newValues = append(newValues, 60, 70, 80, 90, 100)
	fmt.Println(newValues)

	//make is used to create a slice, map or channel only and not an array

	sort.Ints(newValues) // we can sort a slice
	fmt.Println(newValues)

	sort.Sort(sort.Reverse(sort.IntSlice(newValues))) // we can reverse a slice
	fmt.Println(newValues)

	// to remove an element from a slice based on index

	var fruits = []string{"apple", "banana", "mango", "orange", "grape"}
	fmt.Println(fruits)
	var index int = 2
	fruits = append(fruits[:index], fruits[index+1:]...) // we can remove an element from a slice by appending the slice before and after the element to be removed
	fmt.Println(fruits)
}

// Description of Slices

// Slices are a dynamic array. Slices are similar to arrays, but their size is not fixed.
// Slices are used to store a collection of elements of the same type. Slices are more flexible than arrays.
//  Slices are used to implement dynamic arrays. Slices are reference types. Slices are created using the make function.
//   Slices are used to implement stacks, queues, and dynamic lists. Slices are used to implement matrices, vectors, and multi-dimensional arrays.
//    Slices are used to implement strings. Slices are used to implement sparse arrays. Slices are used to implement graphs and trees etc..
