package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Give me rating between 1 and 10")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // 64 is the bit size of the float we want to parse and TrimSpace removes any leading or trailing spaces

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

}
