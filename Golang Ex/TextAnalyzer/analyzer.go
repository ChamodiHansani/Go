package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a text string: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	wordCount := countWords(text)
	charCount := countCharacters(text)
	letterCount := countSpecificLetter(text, 'a') // Example: count occurrences of letter 'a'

	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Number of characters: %d\n", charCount)
	fmt.Printf("Number of occurrences of letter 'a': %d\n", letterCount)
}

// countWords counts the number of words in a given text string.
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

// countCharacters counts the number of characters in a given text string.
func countCharacters(text string) int {
	count := 0
	for _, char := range text {
		if !unicode.IsSpace(char) { // Ignore spaces
			count++
		}
	}
	return count
}

// countSpecificLetter counts the occurrences of a specific letter in a given text string.
func countSpecificLetter(text string, letter rune) int {
	count := 0
	for _, char := range text {
		if char == letter {
			count++
		}
	}
	return count
}
