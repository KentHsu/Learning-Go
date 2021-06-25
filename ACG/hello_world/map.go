package main

import "fmt"

func main() {
	order := map[string]int{
		"Kent":   1,
		"Stone":  2,
		"Alex":   3,
		"Johnny": 4,
	}
	fmt.Println(order)

	reverse := map[string]int{}
	reverse["Kent"] = 4
	reverse["Stone"] = 3
	reverse["Alex"] = 2
	reverse["Johnny"] = 1
	fmt.Println(reverse)

	delete(reverse, "Johnny")
	fmt.Println(reverse)
}
