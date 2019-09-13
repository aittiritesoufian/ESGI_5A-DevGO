package main

import "fmt"

type letter int

const (
	// TODO: use itoa instead of this current declaration.
	a letter = iota
	b
	c
	d
)

var letterList = []string{
	// TODO: Use the const in the key instead of value number
	// and add an unknown value for the first element of the slice.
	a: "a",
	b: "b",
	c: "c",
	d: "d",
}

// TODO add the string method on the letter type that implement the stringer interface.

func main() {
	fmt.Println(letterList) // should output: {<a>, <b>, <c>, <d>}
}
