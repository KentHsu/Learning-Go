package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 18
	ages["Keith"] = 28
	ages["James"] = 67
	ages["Michael"] = 16
	ages["Leigha"] = 5

	// Sequence loop
	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17:
			fmt.Println(name, "has a prime number age")
		case 16:
			fmt.Println(name, "can drive")
		case 18:
			fmt.Println(name, "can vote")
		case 67:
			fmt.Println(name, "can retire")
		default:
			fmt.Println(fmt.Sprintf("There's nothing special about %s's current age.", name))
		}
	}

	// Traditional loop
	fmt.Println("\n")
	for i := 1; i <= 10; i++ {
		fmt.Println("Traditional loop: ", i)
	}

	// Conditional loop
	fmt.Println("\n")
	j := 0
	for j < 10 {
		fmt.Println("Conditonal loop: ", j)
		j++
	}

	// Stopping iteration
	fmt.Println("\n")
	k := 0
	for k < 10 {
		if k%2 == 0 {
			k++
			continue
		} else if k == 5 {
			break
		}
		fmt.Println("Stopping loop:", k)
		k++
	}
}
