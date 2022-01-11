package main

import "fmt"

// dead simple struct
type rect struct {
	width, height int
}

// declaring method fo rect struct
// struct method can be two types
// accepts pointer to the struct as paramter
func (r *rect) area() int {
	return r.width * r.height
}

// as value of the parameter
func (r rect) perimeter() int {
	return 2*r.width + 2 *r.height
}

func playMethods() {
	// initiating struct
	r := rect{width: 10, height: 5}

	// invoking struct methods
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	// dereferncing struct instance r
	// and use it as it is
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())


}

// func main() {
// 	playMethods()
// }