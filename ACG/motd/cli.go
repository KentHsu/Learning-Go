package main

import (
	"bufio"
	"flag"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	// Define flag
	var name string
	var greeting string
	var prompt bool
	var preview bool

	// Parse flag
	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "phrase to use within the message")
	flag.BoolVar(&prompt, "prompt", false, "use prompt to input name and greeting")
	flag.BoolVar(&preview, "preview", false, "use preview to output message without writing to file")

	flag.Parse()

	// Show usage if flags are invalid
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	// Optionally print flags and exit based on DEBUG env variable
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name:", name)
		fmt.Println("Greeting", greeting)
		fmt.Println("Prompt", prompt)
		fmt.Println("Preview", preview)
	}

	// Conditionally read from stdin
	if prompt {
		name, greeting = renderPrompt()
	}

	// Generate message
	message := message.Greeting(name, greeting)

	// Either preview the message or write to file
	if preview {
		fmt.Println(message)
	} else {
		f, err := os.OpenFile("./motd", os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Error: Unable to open the file: ", err)
			os.Exit(1)
		}

		defer f.Close()

		_, err = f.Write([]byte(message))

		if err != nil {
			fmt.Println("Error: Failed to write to the file: ", err)
			os.Exit(1)
		}
	}
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)
	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return
}