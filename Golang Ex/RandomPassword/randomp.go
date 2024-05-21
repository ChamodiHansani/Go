package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func main() {
	var length int
	var includeUpper, includeLower, includeDigits, includeSpecial bool

	// Ask user for password length and character types
	fmt.Print("Enter the desired password length: ")
	fmt.Scan(&length)

	fmt.Print("Include uppercase letters? (true/false): ")
	fmt.Scan(&includeUpper)

	fmt.Print("Include lowercase letters? (true/false): ")
	fmt.Scan(&includeLower)

	fmt.Print("Include digits? (true/false): ")
	fmt.Scan(&includeDigits)

	fmt.Print("Include special characters? (true/false): ")
	fmt.Scan(&includeSpecial)

	password := generatePassword(length, includeUpper, includeLower, includeDigits, includeSpecial)
	fmt.Printf("Generated password: %s\n", password)
}

func generatePassword(length int, includeUpper, includeLower, includeDigits, includeSpecial bool) string {
	var charPool string
	// Construct character pool based on user preferences
	if includeUpper {
		charPool += upperLetters
	}
	if includeLower {
		charPool += lowerLetters
	}
	if includeDigits {
		charPool += digits
	}
	if includeSpecial {
		charPool += specialChars
	}

	// Check if no character types are selected
	if len(charPool) == 0 {
		fmt.Println("No character types selected. Please select at least one character type.")
		return ""
	}

	// Generate password using characters from the character pool
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charPool))))
		if err != nil {
			fmt.Println("Error generating random number:", err)
			return ""
		}
		password[i] = charPool[index.Int64()]
	}

	return string(password)
}
