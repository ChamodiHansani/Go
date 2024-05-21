package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader to read from standard input

	// Ask the user to enter the principal amount
	fmt.Print("Enter the principal amount: ")
	principalStr, _ := reader.ReadString('\n')
	principalStr = strings.TrimSpace(principalStr)
	principal, err := strconv.ParseFloat(principalStr, 64) // Convert the input string to a float64
	if err != nil {
		fmt.Println("Invalid principal amount. Please enter a valid number.")
		return
	}

	// Ask the user to enter the annual interest rate
	fmt.Print("Enter the annual interest rate (in percentage): ")
	rateStr, _ := reader.ReadString('\n')
	rateStr = strings.TrimSpace(rateStr)
	rate, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		fmt.Println("Invalid interest rate. Please enter a valid number.")
		return
	}

	// Ask the user to enter the loan term in months
	fmt.Print("Enter the loan term (in months): ")
	termStr, _ := reader.ReadString('\n')
	termStr = strings.TrimSpace(termStr)
	term, err := strconv.Atoi(termStr)
	if err != nil {
		fmt.Println("Invalid loan term. Please enter a valid number.")
		return
	}

	// Calculate the monthly payment
	monthlyPayment := calculateMonthlyPayment(principal, rate, term)
	// Print the monthly payment
	fmt.Printf("Your monthly payment is: %.2f\n", monthlyPayment) //.2 means that the number will be rounded and displayed with 2 decimal places.
}

// calculateMonthlyPayment calculates the monthly payment for a loan
// using the principal amount, annual interest rate, and loan term in months
func calculateMonthlyPayment(principal float64, annualRate float64, months int) float64 {
	monthlyRate := annualRate / 12 / 100                                // Convert annual interest rate to monthly interest rate
	numerator := monthlyRate * math.Pow(1+monthlyRate, float64(months)) // Calculate the numerator of the formula
	denominator := math.Pow(1+monthlyRate, float64(months)) - 1         // Calculate the denominator of the formula
	return principal * numerator / denominator                          // Calculate and return the monthly payment
}
