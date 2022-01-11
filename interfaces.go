package main

import (
	"fmt"
	"math"
)

// interface is a collection of method signature
// what is method signature anyway?
// it's like describing parameters, return types, name of a method for now, and code it later
// here we defined an interface geometry
// and two method signature in it
type geometry interface {
	area() float64 // area() takes no parameter and returns a float64 value
	perimeter() float64 // likewise
	greetMe(name string) string // this one accepts a string paramter and return string value
}

// a struct named rect with two value
type rect struct {
	width, height float64
}

// likewise
type circle struct {
	radius float64
}

// implementing interface method area for 
// struct named rect
func (r rect) area() float64 {
	return r.width * r.height
}

// implementing interface method perimeter for rect
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// implementing interface method greetMe for rect
func (r rect) greetMe(name string) string {
	return "Hello " + name
}

// implementation of geometry interface for circle struct
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) greetMe(name string) string {
	return "Hello " + name
}

// a generic function accepts geometry type and string type as parameter
func measure(g geometry, name string) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
	fmt.Println(g.greetMe(name))
}

// here's the main sauce!
func playInterface() {
	// declaring instance of rect struct
	r := rect{width: 3, height: 4}
	// declaring instance of circle struct
	c := circle{radius: 5}

	// passing instances in measure function
	// wait!
	// measure accepts geometry type as first parameter, but
	// we are passing struct type, why it's not throwing error?
	// because interface is both collection of methods and type
	// any type that uses methods defined in interface, that type will be automatically determined
	// here both circle and rect uses methods, defined in geometry and it's automatically determined
	measure(r, "rectangle")
	measure(c, "circle")
}

// func main() {
// 	playInterface()
// }