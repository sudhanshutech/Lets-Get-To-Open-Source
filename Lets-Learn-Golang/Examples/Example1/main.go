// This is example 1 that is using struct to define few parameters of user
// and then user give input of there details and we print them there details

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to details page")

	IdReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Id: ")
	Idinput, _ := IdReader.ReadString('\n')

	FirstNameReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your firstname: ")
	FirstNameinput, _ := FirstNameReader.ReadString('\n')

	LastNameReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your lastname: ")
	LastNameinput, _ := LastNameReader.ReadString('\n')

	EmailReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Email: ")
	Emailinput, _ := EmailReader.ReadString('\n')

	StatusReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your status: ")
	Statusinput, _ := StatusReader.ReadString('\n')

	user := UserDetails{ID: Idinput, FirstName: FirstNameinput, LastName: LastNameinput, Email: Emailinput, status: Statusinput}

	fmt.Printf("Your details are: \n")
	fmt.Printf("Name is %v %v", user.FirstName, user.LastName)
	fmt.Printf("Id is %v\n", user.ID)
	fmt.Printf("Email is %v\n", user.Email)
	fmt.Printf("Name is %v\n", user.status)

}

type UserDetails struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	status    string
}
