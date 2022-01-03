package main

import "fmt"

// what is closure?
// in brief, it persists local variable outside of its scope
// in details, check this out "https://stackoverflow.com/questions/36636/what-is-a-closure/7464475"

func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextVal := closures()

	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())


}