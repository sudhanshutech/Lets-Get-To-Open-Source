package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello my time")

	presentTime := time.Now() // this is the current time
	// fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // this date is the reference date for Go to format dates

	createdDate := time.Date(2018, time.December, 29, 23, 59, 59, 0, time.UTC) //this is the date that we want to format and print
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
