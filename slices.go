package main

import "fmt"

func slices() {

	// create a new slice
	// unlike arrays, slice indexes don't have zero value at intialization
	s := make([]string, 4)
	fmt.Println("initialized s: ", s) // print "initialized s: [ ]"

	// set values in slice
	s[0] = "A"
	s[1] = "z"
	s[2] = "a"
	s[3] = "d"

	// get all values
	fmt.Println("get all values:", s) // print "get all values: [A z a d]"
	// get specific index value
	// if tried to access index out of bound,
	// array gives out of bound error instantly
	// whereas slice doesn't give any "out of bound error"
	// it yields error in runtime
	fmt.Println("single value: ", s[0]) // print "single value:  A"

	// length can be printed like array
	fmt.Println("Length of s: ", len(s)) // print "Length of s:  4"

	// slices can be copied into another slice
	c := make([]string, len(s))
	fmt.Println("Value of c before copying: ", c) // print "Value of c before copying:  [   ]"
	copy(c, s)
	fmt.Println("Value of c after copying", c) // print "Value of c after copying [A z a d]"

	// what if we allocate a slice more than the length of "s"?
	// let's see
	d := make([]string, 10)
	fmt.Println("Value of d before copying s:", d) // print "Value of d before copying s: [         ]"
	copy(d, s)
	fmt.Println("Value of d after copying s:", d) // print "Value of d after copying s: [A z a d      ]"
	// turns out it occupies len(s) indexes and left other empty, what a waste!

	// let's do some appending
	// here, 10 index is allocated for d
	// it will append below values after 10 index
	d = append(d, " ")
	d = append(d, "M")
	d = append(d, "a")
	d = append(d, "m")
	d = append(d, "u")
	d = append(d, "n")
	// let's print it
	fmt.Println("Appended d: ", d) // prints "Appended d:  [A z a d         M a m u n]"
	// what's your length?
	// array doesn't have append method
	// append extends slice index length and add values there
	fmt.Println("Current length of d: ", len(d)) // print "Current length of d:  16"

	// slice the slice!
	// lower bound inclusive
	// upper bound exclusive,
	// smells like python
	l := s[1: 4] 
	fmt.Println("l: ", l) // print "l:  [z a d]"

	l = s[:2] // slices up to 2(2 excluded)
	fmt.Println("l2: ", l) // print "l5:  [A z]"

	l = s[2:]
	fmt.Println("l again: ", l) // print "l again:  [a d]"

	// initialize and declare slice
	t := []string{"A", "z", "a", "d"}
	fmt.Println("shorthand slice: ", t) // print "shorthand slice:  [A z a d]"

	// make 2D slice
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // making dynamic row on the fly
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D: ", twoD) // print "2D:  [[0] [1 2] [2 3 4]]"
}

// func main() {
// 	slices()
// }