package strprocess

import (
	"strings"
)

// CustomSplit function customizes the splitting of text into words, treating newline characters separately
func CustomSplit(s string) []string {
	var fields []string
	var currentField string

	// Iterate through each character in the input string
	for _, ch := range s {
		if ch == ' ' || ch == '\t' {
			// Found a whitespace character, add the current field to the slice
			if currentField != "" {
				fields = append(fields, currentField)
				currentField = ""
			}
		} else if ch == '\n' {
			// Treat newline character as a separate field
			if currentField != "" {
				fields = append(fields, currentField)
				currentField = ""
			}
			fields = append(fields, "\n")
		} else {
			// Non-whitespace character, add it to the current field
			currentField += string(ch)
		}
	}

	// Add the last field (if any)
	if currentField != "" {
		fields = append(fields, currentField)
	}

	return fields
}

// CustomJoin function customizes the joining of words into a single string, handling newline characters
func CustomJoin(s []string) string {
	var result strings.Builder

	// Iterate through each field in the input slice
	for i, field := range s {
		// Skip adding a space if the current field is first or current or previous field is '\n'
		if i == 0 || field == "\n" || s[i-1] == "\n" {
			result.WriteString(field)
		} else {
			result.WriteString(" " + field)
		}
	}

	return result.String()
}