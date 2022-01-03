package main

import "fmt"

func ranges() {

	nums := []int{2,3,4}
	sum := 0

	// range iterates over array of integers for sum
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)


	// first value in range is index
	// second value is the iterables objects value itself
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// range iterates over map as key value
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	// only iterate keys or values too
	for k := range m {
		fmt.Println("key:", k)
	}

	// iterating over string returns unicode value of each character
	for i, c := range "Azad" {
		fmt.Println(i, c)
	}
}

// func main() {
// 	ranges()
// }