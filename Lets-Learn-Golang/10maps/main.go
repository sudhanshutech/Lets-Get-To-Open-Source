package main

import "fmt"

func main() {

	languages := make(map[string]string) // we can create a map with make where first string is the key and second string is the value

	languages["en"] = "English"
	languages["fr"] = "French"
	languages["es"] = "Spanish"
	languages["de"] = "German"
	languages["hi"] = "Hindi"

	fmt.Println(languages)
	fmt.Println("en is ", languages["en"])
	fmt.Println("fr is ", languages["fr"])
}
