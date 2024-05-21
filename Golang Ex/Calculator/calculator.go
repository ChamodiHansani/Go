package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader to read from standard input

	// Ask the user for the first number
	fmt.Print("Enter the first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1, err := strconv.ParseFloat(strings.TrimSpace(num1Str), 64) // Convert the input string to a float64
	if err != nil {
		fmt.Println("Invalid input for the first number. Please enter a valid number.")
		return // Exit the program if the input was invalid
	}

	// Ask the user for the second number
	fmt.Print("Enter the second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(num2Str), 64) // Convert the input string to a float64
	if err != nil {
		fmt.Println("Invalid input for the second number. Please enter a valid number.")
		return
	}

	// Prompt the user for the operation
	fmt.Print("Enter the operation (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	// Calculate the result based on the operation
	result, err := calculate(num1, num2, op)
	if err != nil { // Check if there was an error during the calculation
		fmt.Println("Error:", err)
		return // Exit the program if there was an error
	}

	// Print the result
	fmt.Printf("The result of %.2f %s %.2f is: %.2f\n", num1, op, num2, result)
}

// calculate performs the specified operation on the two numbers
func calculate(num1, num2 float64, op string) (float64, error) {
	switch op {
	case "+": // Addition
		return num1 + num2, nil
	case "-": // Subtraction
		return num1 - num2, nil
	case "*": // Multiplication
		return num1 * num2, nil
	case "/": // Division
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return num1 / num2, nil
	default: // Invalid operation
		return 0, fmt.Errorf("invalid operation")
	}
}
