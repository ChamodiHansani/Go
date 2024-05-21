package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Input and output file names
	inputFileName := "input.txt"
	outputFileName := "output.txt"
	wordToFind := "Go"

	// Read data from the input file
	lines, err := readLinesFromFile(inputFileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Process data
	lineCount := len(lines)
	wordCount := countWordOccurrences(lines, wordToFind)

	// Prepare results
	results := fmt.Sprintf("Total lines: %d\nOccurrences of the word '%s': %d\n", lineCount, wordToFind, wordCount)

	// Write results to the output file
	err = writeToFile(outputFileName, results)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
	} else {
		fmt.Println("Results written to", outputFileName)
	}
}

// readLinesFromFile reads a file and returns its lines as a slice of strings
func readLinesFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// countWordOccurrences counts the number of occurrences of a specific word in a slice of lines
func countWordOccurrences(lines []string, word string) int {
	count := 0
	for _, line := range lines {
		count += strings.Count(line, word)
	}
	return count
}

// writeToFile writes a string to a file
func writeToFile(fileName string, data string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}
