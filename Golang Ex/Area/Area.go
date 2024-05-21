package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask the user for the length
	fmt.Print("Enter the length of the rectangle: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	length, err := strconv.ParseFloat(lengthInput, 64)
	if err != nil {
		fmt.Println("Invalid input for length. Please enter a valid number.")
		return
	}

	// Ask the user for the width
	fmt.Print("Enter the width of the rectangle: ")
	widthInput, _ := reader.ReadString('\n')
	widthInput = strings.TrimSpace(widthInput)
	width, err := strconv.ParseFloat(widthInput, 64)
	if err != nil {
		fmt.Println("Invalid input for width. Please enter a valid number.")
		return
	}

	// Calculate the area
	area := length * width

	// Print the result
	fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
