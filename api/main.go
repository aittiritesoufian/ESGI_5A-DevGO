package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	now := time.Now()
	u := User{
		FirstName:   "Henri",
		LastName:    "Lepic",
		DateOfBirth: now,
	}
	fmt.Println(u)
	//payload, err := json.Marshal(u)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(payload))

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, u)
	})
	r.Run(":8080")
}

type User struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"birth_date"`
}

func (u User) String() string {
	return u.FirstName + " " + u.LastName
}
