package main

import (
	"fmt"
	"motd/message"
)

func main() {
	// function definition 1
	greet := greeting("Kent", "Hello")
	fmt.Println(greet)

	// function definition 2
	fmt.Println(greeting2("Stone", "Hello"))

	// exposed function
	exposed_message := message.Greeting("Alex", "Hello")
	fmt.Println(exposed_message)
}

func greeting(name string, message string) string {
	return fmt.Sprintf("%s, %s", message, name)
}

func greeting2(name, message string) (salutation string) {
	salutation = fmt.Sprintf("%s, %s", message, name)
	return
}
