package main

// dead simple recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n - 1)
}


// func main() {
// 	fmt.Println(fact(10))

// 	// closure can be recursive
// 	// it has to be declared before with typed var
// 	var fib func(n int) int

// 	// and then store the result of annonymous function in fib
// 	fib = func(n int) int {
// 		if n < 2 {
// 			return n
// 		}
// 		return fib(n-1) + fib(n-2)
// 	}

// 	fmt.Println(fib(3))
// }