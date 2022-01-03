package main

import "fmt"

func sum(a, b int) int {

	return a + b;

}

// return array
func array() []int {
	a := []int{1, 2, 3}

	return a
}

// return multiple values
func multiples() (int, string) {
	return 1, "one"
}

// variadic function
// what is it?
// it simply takes arbitrary number of arguments of specific type
func nSums(nums ...int) int {
	// this function takes arbitrary number of integers
	fmt.Print(nums, " ") // Print "[1 2 3 4]", Print iterates over iterable data structure
	total := 0

	for _, num := range nums {
		total += num
	}

	return total 
}



// func main() {
// 	res := sum(10, 20)

// 	fmt.Println(res)

// 	a := array()
// 	fmt.Println(a) // print "[1 2 3]"

// 	i, s := multiples()
// 	fmt.Println("multiple returns:", i, s) // print "multiple returns: 1 one"

// 	totalSum := nSums(1, 2, 3, 4) // print "nSums: 10"
// 	fmt.Println("\nnSums:", totalSum)

// 	// pass array to nSums
// 	nums := []int{1,2,3,4,10}
// 	fmt.Println("\n",nSums(nums...)) // print "20"
// }