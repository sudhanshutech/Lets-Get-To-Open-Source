package main

import "fmt"

func main() {

	user := User{ID: 1, FirstName: "John", LastName: "Doe", Email: "email@gmail.com", Status: true}
	fmt.Println(user)
	fmt.Printf("User details are: %+v\n", user)                                            // %+v will print the field names as well
	fmt.Printf("Name is %v %v and email is %v", user.FirstName, user.LastName, user.Email) // %v will print the values only
}

// User is a struct and struct is a collection of fields
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Status    bool
}

// Description of Structs

// Structs are user-defined types that are used to group related data. Structs are similar to classes in object-oriented programming.
