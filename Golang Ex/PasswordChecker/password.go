package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	strength := checkPasswordStrength(password)
	fmt.Printf("Password strength: %s\n", strength)
}

func checkPasswordStrength(password string) string {
	var (
		lengthValid, uppercaseValid, lowercaseValid, numberValid bool
		length, uppercaseCount, lowercaseCount, numberCount      int
	)

	// Check length
	length = len(password)
	lengthValid = length >= 8 && length <= 20

	// Check uppercase, lowercase, and numbers
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			uppercaseCount++
		case unicode.IsLower(char):
			lowercaseCount++
		case unicode.IsNumber(char):
			numberCount++
		}
	}

	uppercaseValid = uppercaseCount > 0
	lowercaseValid = lowercaseCount > 0
	numberValid = numberCount > 0

	// Determine strength based on criteria
	if lengthValid && uppercaseValid && lowercaseValid && numberValid {
		return "Strong"
	} else if lengthValid && (uppercaseValid || lowercaseValid || numberValid) {
		return "Moderate"
	} else {
		return "Weak"
	}
}
