package main

import (
	"fmt"
	"time"
)

// goroutines call functions asynchronously
// to make a function in goroutine just add go before function

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func playRoutine() {
	// invoking synchronous function call
	f("direct")

	// invoking function with goroutine
	go f("goroutine")

	// invoking annonymous function with goroutine
	go func (msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

// above function first call synchronous functions f("direct")
// while running other two function(goroutine) in separate goroutines

func main() {
	playRoutine()
}