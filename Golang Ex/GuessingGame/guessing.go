package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator with the current time
	target := rand.Intn(100) + 1     // Generate a random number between 1 and 100

	reader := bufio.NewReader(os.Stdin) // Create a new reader to read from standard input
	var guess int
	var err error

	fmt.Println("Guess the number (between 1 and 100):")

	// Start an infinite loop to prompt the user for guesses
	for {
		fmt.Print("Enter your guess: ")
		guessStr, _ := reader.ReadString('\n') // Read the input from the user until a newline character is encountered
		guessStr = strings.TrimSpace(guessStr)
		guess, err = strconv.Atoi(guessStr) // Convert the input string to an integer
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		// Check if the guess is within the valid range
		if guess < 1 || guess > 100 {
			fmt.Println("Please enter a number between 1 and 100.")
			continue
		}

		// Compare the guess with the target number
		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d!\n", target)
			break
		}
	}
}
