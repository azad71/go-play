package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

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

func playInterfaces() {
	animals := []Animals{new(Dog), new(Cat), Bird{}, Fish{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

var input = `
	{
		"created_at": "Thu May 31 00:00:01 +0000 2012" 
	}
`

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

func decode_json() {
	var val map[string]interface{}

	if err:= json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)

	for k, v:= range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}

// func main() {
// 	// playInterfaces()
// 	decode_json()

// }