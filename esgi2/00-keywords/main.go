// write a code that can comile and use at least 10 golang keywords.
package main

import (
	"fmt"
	"time"
)

func main() {
	const test = "coucou"
	if true {
		Display(test)
	}
	fmt.Println(Display("Salut a tous"))
}

// Print the given message in the standard output
func Display(msg string) string {
	switch msg {
	case "coucou":
		fmt.Println("Coucou !")
	case "test":
		fmt.Println("Test !!!")
	case "the end":
		fmt.Println("The End !!!!")
	default:
		fmt.Println("It's default!")
	}
	var t = time.Now()
	fmt.Println("The current time is : ")
	fmt.Println(t)
	return "plop"
}