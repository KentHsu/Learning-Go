package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 70

	// if statement
	if ages["Kevin"] > 65 {
		fmt.Println("Kevin is of retired age")
	} else if ages["Kevin"] < 18 {
		fmt.Println("Kevin will start his career soon...")
	} else {
		fmt.Println("Kevin is not of retired age")
	}

	// switch statement
	switch {
	case ages["Kevin"] > 65:
		fmt.Println("Kevin is of retired age")
	case ages["Kevin"] < 18:
		fmt.Println("Kevin will start his career soon...")
	default:
		fmt.Println("Kevin is not of retired age")
	}

	// more useful switch statement
	switch ages["Kevin"] {
	case 1, 2, 3, 5, 7, 11, 13, 17:
		fmt.Println("Kevin is of prime age")
	case 18:
		fmt.Println("Kevin is an adult")
	case 65:
		fmt.Println("Kevin is ready to retired age")
	default:
		fmt.Println("Kevin's age has nothing special")
	}
}
