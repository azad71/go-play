package main

import "fmt"

// what is closure?
// in brief, it persists local variable outside of its scope
// in details, check this out "https://stackoverflow.com/questions/36636/what-is-a-closure/7464475"

// an annonymous function declared inside function closures()
func outer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// nextVal will hold the reference for inner annonymous function
	// after invoking outer(), it will create a variable i
	// initial value of i is 0, as defined in function
	nextVal := outer()

	// everytime nextVal() is called, its internal value will be updated
	// since closure persists local variable(in this case, i), i will hold the updated value
	// smells like javascript generator!
	fmt.Println(nextVal()) // print "1"
	fmt.Println(nextVal()) // print "2"
	fmt.Println(nextVal()) // print "3"
	fmt.Println(nextVal()) // print "4"


	// what if we call outer function again?
	nextVal2 := outer()

	fmt.Println(nextVal2()) // print "1"
	// interesting! 
	// invoking nextVal2(), starts i from 1
	// what will nextVal() yield, if invoke now?
	// let's see!
	fmt.Println(nextVal()) // print "5"

	// interesting! each instance of outer() is unique 
	// and holds their own value
	// that's why nextVal() still keep incrementing even after
	// outer function is invoked in nextVal2()


}