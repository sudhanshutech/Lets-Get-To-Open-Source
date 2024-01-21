package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("This is working with files example")

	content := "This is a test file which is inserted by Go program"

	file, err := os.Create("test.txt") // creates a file
	if err != nil {                    // if there is an error creating the file, throw an error
		fmt.Println("Unable to create file:", err)
		panic(err) // panic is used to stop the execution of program
	}

	length, err := io.WriteString(file, content) // writes the content to the file
	if err != nil {
		fmt.Println("Unable to write to file:", err)
		panic(err)
	}

	fmt.Println("Length of file is :", length)
	defer file.Close()
	readFile("test.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName) // os .ReadFile returns byte array and error and is used to read the file

	// if err != nil {
	// 	fmt.Println("Unable to read file:", err)
	// 	panic(err)
	// }
	checkError(err)

	fmt.Println("File content is:", string(dataByte)) // convert byte to string because ReadFile returns byte array
}

// we are using nil check err so many times, so we can create a function to handle it
func checkError(err error) {
	if err != nil {
		fmt.Println("Error is:", err)
		panic(err)
	}
}
