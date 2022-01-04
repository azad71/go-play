package main

import "fmt"

// golang supports both
// pass by value and
// pass by reference

// here passing a copy of val
// won't change original value of val
// since it's a distinct copy
func passByValue(val int) {
	val = 0
}

// here passing the pointer of ptr memory location
// will change it's original value
// since it's original memory address passed here as parameter
func passByReference(ptr *int) {
	*ptr = 0
}

func playPointers() {
	i := 1

	fmt.Println("Initial value:", i) // print "Initial value: 1"

	// passing the value of i
	passByValue(i)
	fmt.Println("After passByValue:", i) // print "After passByValue: 1"

	// passing the memory address of i
	passByReference(&i)
	fmt.Println("After passByReference:", i) // print "After passByReference: 0"
	// i changed, not me!
	
	// pointer is also printable
	fmt.Println("I(pointer) look like this:", &i)
}

// func main() {
// 	play()
// }