package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// isValidCombo checks if a combo is valid (basic email and password format check)
func isValidCombo(combo string) bool {
	parts := strings.Split(combo, ":")
	if len(parts) != 2 {
		return false
	}

	email := parts[0]
	password := parts[1]

	// Simple regex for email validation
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(email) || len(password) < 6 {
		return false
	}

	return true
}

func main() {
	// Open the file containing combos
	file, err := os.Open("combos.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new file to save the cleaned combos
	outputFile, err := os.Create("cleaned_combos.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		combo := scanner.Text()
		if isValidCombo(combo) {
			outputFile.WriteString(combo + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
