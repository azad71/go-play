package main

import (
	"fmt"
	"math"
)

// constant variables are two types
// 1. typed 2. untyped
// below is typed constant
const s string = "I AM CONSTANT";

func constants() {
	fmt.Println(s)

	// this one is untyped constant
	const n  = 50000000000

	
	const d = 3e20 / n 

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

// func main() {
// 	constants()
// }