package main

import "fmt"

func main() {
	fmt.Println("Interpreted String Literal")
	fmt.Println(`Raw String Literal`)
	fmt.Println(`
This is multi-line. \n
String
`)
	fmt.Println("\u2272")
	fmt.Println('A')
}
