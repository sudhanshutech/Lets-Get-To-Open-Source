// Description: This program will make a web request to the given url and will print the response
//  and also the type of response

package main

import (
	"fmt"
	"net/http"
)

const url = "https://dummy.restapiexample.com/api/v1/employees" // sample url

func main() {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
	fmt.Printf("Response is of type: %T\n", response)

	response.Body.Close() // This is important to close the connection because if we don't close the connection then it will be open and it will be a memory leak
}
