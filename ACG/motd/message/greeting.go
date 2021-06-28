package message

import "fmt"

func Greeting(name string, message string) string {
	return fmt.Sprintf("%s, %s", message, name)
}
