package main

import "fmt"
import "time"
import "github.com/hello/myint"

func main(){
	var mi = myint.MyType(15)
	fmt.Println(mi.Divide(5))
	fmt.Println(mi.Add(5))
	fmt.Println(mi.Sub(5))
	fmt.Println(mi.Multiply(5))
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