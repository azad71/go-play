package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2 *r.height
}

func playMethods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())

	rp.height = 20
	fmt.Println("height: ", rp.height)
}

func main() {
	playMethods()
}