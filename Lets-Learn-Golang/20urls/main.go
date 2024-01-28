package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://api/keychains?limit=100"

func main() {
	fmt.Println("This is about handling urls")

	result, _ := url.Parse(myUrl) // This will parse the url and will give us the result
	fmt.Println(result)

	fmt.Println(result.Scheme)   // This will print the scheme of the url means http or https
	fmt.Println(result.Host)     // This will print the host of the url
	fmt.Println(result.Path)     // This will print the path of the url
	fmt.Println(result.RawQuery) // This will print the query of the url

	queryParams := result.Query() // This will give us the query params of the url
	fmt.Printf("Query params are: %v\n", queryParams)
	fmt.Printf("Limit is: %v\n", queryParams["limit"][0]) // This will print the limit of the url

	// Now constructing the url

	partsOfUrl := &url.URL{ // This will create a new url, here &url.URL is a pointer to the url
		Scheme:   "https",
		Host:     "api",
		Path:     "auth/keys",
		RawQuery: "limit=100",
	}

	newUrl := partsOfUrl.String() // This will convert the url to string
	fmt.Println(newUrl)
}
