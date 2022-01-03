package main

import "fmt"

func vaiables() {
	// declare a variable like this
	// if value assigned when variable is declared
	var a = "hello";
	fmt.Println(a);

	// this is the shorthand of above code
	f := "apple";
	fmt.Println(f);

	// one liner for declaring multiple variables
	var b, c int = 10, 20;
	fmt.Println(b, c)

	// boolean value
	var d = true;
	fmt.Println(d);

	// if no value assigned it will have 0 as default value(for int)
	var e int;
	fmt.Println(e);


}

// func main() {
// 	vaiables()
// }