package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	// ReadString will block until the delimiter is entered
	//' _' means that we don't care about the value returned from ReadString
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", input, "!")
}
