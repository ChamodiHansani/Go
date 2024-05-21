package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader to read from standard input

	fmt.Print("Enter your name: ")     // Ask the user for their name
	name, _ := reader.ReadString('\n') // Read the input from the user until a newline character is encountered
	name = strings.TrimSpace(name)     // Remove any leading or trailing whitespace characters

	fmt.Printf("Hello, %s!\n", name) // Print the greeting with the user's name
}
