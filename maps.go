package main

import (
	"fmt"
)

func maps() {
	// pass map into builtin make method to construct a map
	m := make(map[string]int)

	// classic key value pair
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println("map:", m)

	// extracting values from map
	v1 := m["a"]
	fmt.Println("v1", v1)

	// builtin len() can return the size of map
	fmt.Println("len", len(m))

	// deleting map key
	delete(m, "b")
	fmt.Println("map:", m)

	// prs indicates whether extracted value present in the map or not
	_, prs := m["a"]
	fmt.Println("prs:", prs)

	// map shorthand
	n := map[string]int{"x": 10, "y": 20, "z": 30}
	fmt.Println(n)
}

// func main() {
// 	maps()
// }