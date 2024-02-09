package main

import (
	"os"
	"github.com/01-edu/z01"
)

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func split(input string) []string {
	var result []string
	start := 0

	for i, char := range input {
		if isSpace(char) {
			if i > start {
				result = append(result, input[start:i])
			}
			start = i + 1
		}
	}

	if start < len(input) {
		result = append(result, input[start:])
	}

	return result
}

func rostring(s string) string {
	words := split(s)
	if len(words) < 2 {
		return s
	}

	// Rotate the words to the left
	rotatedWords := append(words[1:], words[0])

	// Join the rotated words into a single string
	var result string
	for i, word := range rotatedWords {
		if i > 0 {
			result += " "
		}
		result += word
	}

	return result
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	output := rostring(input)

	for _, char := range output {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
