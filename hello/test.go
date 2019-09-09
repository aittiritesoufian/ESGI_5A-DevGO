package main

import "fmt"

func main(){
	var test = [109999]string{999:"a"}
	fmt.Println(test[999])
	fmt.Println(len(test))
}