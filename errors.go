package main

import (
	"errors"
	"fmt"
)

// throwing inbuilt error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Something went wrong with 42")
	}

	return arg + 3, nil
}

// type for custom error
type argError struct {
	arg int
	prob string
}

// attaching method on custom error type
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// throwing error with custom error method
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Something went wrong!"}
	}

	return arg + 3, nil
}



func playErrors() {
	// in built error
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	// custom error
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ok, ae)
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func main() {
	playErrors()
}