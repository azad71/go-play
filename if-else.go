package main

import "fmt"

func conditionals() {

	// classic if/else branching
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}


	// declare a variable in if/else on the fly
	// here num variable is only available inside below if/else
	// num can't be accessed outside if/else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// point to note, there's no ternary if in Golang

}

// func main() {
// 	conditionals()
// }