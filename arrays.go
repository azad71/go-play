package main

import "fmt"

func arrays() {
	fmt.Println("array examples: ")

	// initialize an integer array
	// initially elements are zero valued
	var a[5] int 
	fmt.Println("emp: ", a)

	// assign values to index
	a[4] = 100

	fmt.Println("set: ", a)

	// access index
	fmt.Println("get: ", a[4])

	// built in method to get the length of array
	fmt.Println("len: ", len(a))

	// initialize array variable with predefined values
	b := [5]int {1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)


	// shorthand declaration with no size
	temp := [...]int{1,2,3,4,5}
	fmt.Println(temp, len(temp))

	// two dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j:=2; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
 }



// func main() {
// 	arrays()
// }