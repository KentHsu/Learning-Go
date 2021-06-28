package main

import (
	"bufio"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("./motd", os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error: Unable to open the file: ", err)
		os.Exit(1)
	}

	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := message.Greeting(name, phrase)
	_, err = f.Write([]byte(message))

	if err != nil {
		fmt.Println("Error: Failed to write to the file: ", err)
		os.Exit(1)
	}
}
