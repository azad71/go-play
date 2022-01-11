package main

import "fmt"

type Animals interface {
	Speak() string;
}

type Dog struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Bird struct {}

func (b Bird) Speak() string {
	return "Cacao!"
}

type Fish struct {}

func (f Fish) Speak() string {
	return "? ? ? ?"
}

func playInterface() {
	animals := []Animals{new(Dog), new(Cat), Bird{}, Fish{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

func main() {
	playInterface()

}