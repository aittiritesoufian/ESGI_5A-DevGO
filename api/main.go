package main

import "fmt"
import "time"

func main() {
	now := time.Now()
    u := User {
		FirstName:"Henri",
		LastName:"Lepic",
		DateOfBirth:now,
	}
	fmt.Println(u)
}

type User struct{
	FirstName, LastName string
	DateOfBirth time.Time
}