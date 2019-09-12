package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

const (
	NBGoroutines = 100
)

func main() {
	wg.Add(NBGoroutines)
	c := make(chan int)
	for i := 0; i < NBGoroutines; i++ {
		go add(c, i)
	}
	for j := 0; j < NBGoroutines; j++ {
		fmt.Println(<-c)
	}
	wg.Wait()
}

func add(c chan int, value int) {
	c <- value
	wg.Done()
}
