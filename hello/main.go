package main

import "fmt"
import "time"

func main(){
	var a = []string{"A","B","C","D"}
	fmt.Println(a)
	b := make([]string, len(a))
	copy(a,b)
	a[0] = "z"

	fmt.Println(a)
	fmt.Println(b)
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