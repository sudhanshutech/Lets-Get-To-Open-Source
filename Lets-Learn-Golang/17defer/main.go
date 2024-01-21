// Description:
// A defer statement defers the execution of a function until the surrounding function returns.
// Defer is used to ensure that a function call is performed later in a program’s execution,
// usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
// LIFO order (last-in-first-out): the last deferred function is executed first.

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	println("hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
