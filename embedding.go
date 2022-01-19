// embedding allows structs to be embedded into other structs
package main

import "fmt"

// a base struct
type base struct {
	num int
}

// describe method is linked with base struct
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// base struct emebedded inside container
// container will inherit all properties of base
type container struct {
	base
	str string
}

func playEmbedding() {

	// creating struct with literals
	co := container {
		base: base{ // initialize embedded property this way
			num: 1,
		},
		str: "Hemlo, I am Strimg",
	}

	// we can access base property directly inside co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// or type the full path
	fmt.Println("Alson num:", co.base.num)

	// since container inherited all the properties of base
	// and describe method attached to base struct
	// container instance co can now access describe method too
	fmt.Println("describe:", co.describe())

	// I don't understand this part clearly
	// may be dig later, get going now...
	// I get the idea of embedding in golang anyway/btw
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}

func main() {
	playEmbedding()
}