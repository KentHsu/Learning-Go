package main

import "fmt"

func main() {

	// arrays declaration
	//var names = [4]string{"Kent", "Stone", "Alex", "Johnny"}
	names := [4]string{"Kent", "Stone", "Alex", "Johnny"}

	// slics declaration
	names_slices := []string{}
	names_slices = append(names_slices, "Kent")
	names_slices = append(names_slices, "Stone", "Alex", "Johnny")

	// using `make` function
	make_names := make([]string, 4)
	make_names[0] = "Kent"
	make_names[1] = "Stone"
	make_names[2] = "Alex"
	make_names[3] = "Johnny"

	fmt.Println(names)
	fmt.Println(names_slices)
	fmt.Println(make_names)
}
