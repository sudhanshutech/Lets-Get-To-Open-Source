package main // this is a package declaration and every go program should start with a package declaration

import "fmt" // this is an import statement and we are importing the fmt package which is used for formatting

func main() { // this is a main function and every go program should have a main function, it is the entry point for the program
	fmt.Println("Hello, world!")
}

//run GOOS="linux" for linux and this commandis used to compile the code for other OS
