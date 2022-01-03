package main

import (
	"fmt"
	"time"
)

func switchExamples() {
	i := 2

	fmt.Println("Write ", i, " as ")

	// classic switch syntax
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// case with multiple expressions
	switch time.Now().Weekday() {
	case time.Saturday, time.Friday:
		fmt.Println("It's the weekend, Yay!")
	default:
		fmt.Println("It's a working day :(")
	}

	t := time.Now()

	// another way of writing if/else conditions
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// assign function to variable
	// which have switch case inside it
	tellMyType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("Don't know your type", t)
		}
	}

	tellMyType(true)
	tellMyType(1)
	tellMyType("hey")
}

// func main() {
// 	switchExamples()
// }