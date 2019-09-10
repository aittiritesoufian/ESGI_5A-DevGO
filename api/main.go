package main

import (
	"fmt"
	"log"
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
	r.POST("user", func(c *gin.Context) {
		log.Println(c.Request)
		c.JSON(200, nil)
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
