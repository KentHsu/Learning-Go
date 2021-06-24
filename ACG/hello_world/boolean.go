package main

import "fmt"

func main() {
	fmt.Println("This is true: ", true)
	fmt.Println("This is also true: ", 1 < 2)
	fmt.Println("This is false: ", false)
	fmt.Println("This is also false: ", 1 > 2)
	fmt.Println("This is nothing in Go: ", nil)
}
