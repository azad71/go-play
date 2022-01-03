package main

import "fmt"

// C style struct
// what is struct?
// struct group together multiple variables into a recored

// example of a struct
type person struct {
	name string
	age int
}

// this function constructs a person struct
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 26

	// local variables in this scope will survive in global scope
	// returning pointer of local variable
	return &p 
}

func main() {
	fmt.Println(person{"Jon", 26})
	fmt.Println(person{name: "Snow", age: 25})
	fmt.Println(person{name: "Hodor"})
	fmt.Println(&person{name: "Sansa Stark", age: 22})
	fmt.Println(newPerson("Jaime"))

	n := person{name: "Arya", age: 16}
	fmt.Println(n.name, n.age)

	np := &n

	fmt.Println(np.name, np.age)

	np.age = 18
	fmt.Println(np.name, np.age)


	x := newPerson("Jaime")
	fmt.Println(x.name, x.age)
	x.age = 35

	fmt.Println(x)
}