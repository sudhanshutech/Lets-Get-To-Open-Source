package main

import "fmt"

func main() {

	user := User{ID: 1, FirstName: "John", LastName: "Doe", Email: "email@gmail.com", Status: true}
	fmt.Println(user)
	fmt.Printf("User details are: %+v\n", user)                                              // %+v will print the field names as well
	fmt.Printf("Name is %v %v and email is %v\n", user.FirstName, user.LastName, user.Email) // %v will print the values only

	user.GetStatus() // calling a method
	user.AlterEmail()
}

// User is a struct and struct is a collection of fields
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Status    bool
}

// Methods

func (u User) GetStatus() {
	fmt.Println("This is GetStatus method and status is ", u.Status)
}

func (u User) AlterEmail() {
	u.Email = "hr@co.in"
	fmt.Println("This is AlterEmail method and email is ", u.Email)
}

// Description of Methods

// Methods are functions that are associated with a particular type. Methods are similar to functions but they are defined with a receiver.
// The receiver can be a struct type or a non-struct type. The receiver is like a parameter to the method. The receiver can be accessed inside the method.
// The receiver can be a pointer or a non-pointer type. If the receiver is a pointer type then the value passed to the method will be a pointer to the type.
// If the receiver is a non-pointer type then the value passed to the method will be a copy of the type. The receiver can be accessed inside the method.
