package main

import "fmt"

func forLoop() {
	i := 1;

	// while flavored loop
	// declare counter variable outside
	// keep only break conditions in the loop
	// increment/decrement counter variable inside loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// proper for loop
	// counter variable; break condition; increment/decrement
	for j := 1; j <=  9; j++ {
		fmt.Println(j)
	}

	// infinite loop
	for {
		fmt.Println("LOOP")
		break
	}

	// continue keyword skips to next iteration
	for n := 0; n <= 5; n++ {
		// equality comparison comes with ==
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// reverse loop
	for n := 10; n > 0; n-- {
		fmt.Println(n)
	}
}

// func main() {
// 	forLoop()
// }