package main

import "fmt"
import "time"

type myln int32

func main(){
	var mi = myln(15)
	fmt.Println(mi.Divide(5))
	fmt.Println(mi.Add(5))
	fmt.Println(mi.Sub(5))
	fmt.Println(mi.Multiply(5))
}

func (mi myln) Divide(n int) myln{
	return mi / myln(n)
}
func (mi myln) Add(n int) myln{
	return mi + myln(n)
}
func (mi myln) Sub(n int) myln{
	return mi - myln(n)
}
func (mi myln) Multiply(n int) myln{
	return mi * myln(n)
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