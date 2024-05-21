package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// celsiusToFahrenheit converts a temperature from Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// fahrenheitToCelsius converts a temperature from Fahrenheit to Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader to read from standard input

	fmt.Println("Temperature Converter")
	fmt.Println("1: Celsius to Fahrenheit")
	fmt.Println("2: Fahrenheit to Celsius")
	fmt.Print("Choose the conversion (1 or 2): ")

	// Read the user's choice
	choiceStr, _ := reader.ReadString('\n') // Read the input from the user
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr) // Convert the input string to an integer
	if err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Invalid choice. Please enter 1 or 2.")
		return
	}

	// Ask the user to enter the temperature to convert
	fmt.Print("Enter the temperature to convert: ")
	tempStr, _ := reader.ReadString('\n')
	tempStr = strings.TrimSpace(tempStr)
	temp, err := strconv.ParseFloat(tempStr, 64) // Convert the input string to a float64
	if err != nil {
		fmt.Println("Invalid temperature. Please enter a valid number.")
		return
	}

	// Temperature conversion based on the user's choice
	if choice == 1 { // Celsius to Fahrenheit
		convertedTemp := celsiusToFahrenheit(temp)
		fmt.Printf("%.2f°C is %.2f°F\n", temp, convertedTemp) // Print the converted temperature
	} else { // Fahrenheit to Celsius
		convertedTemp := fahrenheitToCelsius(temp)
		fmt.Printf("%.2f°F is %.2f°C\n", temp, convertedTemp) // Print the converted temperature
	}
}
