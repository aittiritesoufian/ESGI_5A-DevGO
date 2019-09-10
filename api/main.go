package main

import "fmt"
import "time"
import "encoding/json"

func main() {
	now := time.Now()
    u := User {
		FirstName:"Henri",
		LastName:"Lepic",
		DateOfBirth:now,
	}
	fmt.Println(u)
	payload,err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(payload))
}

type User struct{
	FirstName string `json:"first_name"`
	LastName string
	DateOfBirth time.Time
}

func (u User) String() string {
	return u.FirstName + " " + u.LastName
}