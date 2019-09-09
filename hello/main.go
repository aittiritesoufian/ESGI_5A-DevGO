package main

import "fmt"
import "time"

func main(){
	var test = [110000]string{999:"a"}
	fmt.Println(test[999])
	fmt.Println(len(test))
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