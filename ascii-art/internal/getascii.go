package internal

import "fmt"

func GetAsciiArt(input string) {
	fontArray := GetBanners() // Read ASCII art font from standard.txt
	// Return if read font has been failed
	if fontArray == nil {
		return
	}
	var result []string
	line := 0 // Initialize line count
	
	for index, ch := range input {
		if ch == '\n' { // Check if character is newline
			line = len(result) // Update line count
			if index == 0 {
				result = append(result, "") // Add empty line if newline is the first character
				line = 1
			}
			if index == len(input)-1 {
				result = append(result, "") // Add empty line if newline is the last character
			} else if input[index+1] == '\n' {
				result = append(result, "") // Add empty line if next character is also newline
			}
		} else if ValidAscii(ch) { //Check if valid ASCII character passed
			if len(result) < line+8 { // Ensure result array has enough space for 8 lines per character
				for j := 0; j < 8; j++ {
					result = append(result, "")
				}
			}
			for i := line; i < line+8; i++ { // Add ASCII art lines for the character
				result[i] += fontArray[ch-32][i-line]
			}
		}
		if !ValidAscii(ch) && ch!='\n' { // Return if ASCII character is not valid
			fmt.Println("Incorrect input text, characters are out of printable ASCII symbols")
			return
		}
	}
	PrintAscii(result) // Print the generated ASCII art
}

// Checks if character is printable ASCII
func ValidAscii(ch rune) bool {
	return ch>=32 && ch<=126
}
