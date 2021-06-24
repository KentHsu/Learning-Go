package main

import "fmt"

func main() {

	// pre-declare variable
	var name string
	name = "Kent"

	// variable shorthand
	// var myInt int = 16
	myInt := 16

	// using blank identifier if the variable is not used
	// var val, ok = "yes", true
	var val, _ = "yes", true

	fmt.Println("name is:", name)
	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is:", val)
	// fmt.Println("ok is:", ok)
}
