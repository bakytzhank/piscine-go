package processors

import (
	"strings"
	"go-reloaded/convert"
	"go-reloaded/strprocess"
)

// ProcessTriggers function handles specific triggers in the text and modifies the text accordingly
func ProcessTriggers(index int, val string, splitTxt []string) {
	switch val {
	case "(up)":
		splitTxt[index-1] = strings.ToUpper(splitTxt[index-1])
		splitTxt[index] = ""
	case "(up,":
		convert.ProcessCase(index, splitTxt, strings.ToUpper)
	case "(low)":
		splitTxt[index-1] = strings.ToLower(splitTxt[index-1])
		splitTxt[index] = ""
	case "(low,":
		convert.ProcessCase(index, splitTxt, strings.ToLower)
	case "(cap)":
		splitTxt[index-1] = strings.Title(strings.ToLower(splitTxt[index-1]))
		splitTxt[index] = ""
	case "(cap,":
		convert.ProcessCase(index, splitTxt, func(s string) string {
			return strings.Title(strings.ToLower(s))
		})
	case "(hex)":
		convert.ConvertToDecimal(index, splitTxt, 16)
	case "(bin)":
		convert.ConvertToDecimal(index, splitTxt, 2)
	case "A":
		convert.ConvertAtoAn(index, splitTxt)
	case "a":
		convert.ConvertAtoAn(index, splitTxt)
	}
}

// ProcessPunctuation function handles punctuation in the processed text
func ProcessPunctuation(splitTxt []string) []byte {
	text := strprocess.CustomJoin(splitTxt)
	bytes := []byte(text)

	// Iterate through each character in the text
	for i := 1; i <= len(bytes); i++ {
		for i, c := range bytes {
			// Check for specific punctuation characters and swap places with preceding space
			if c == '.' || c == ',' || c == '!' || c == '?' || c == ';' || c == ':' {
				if bytes[i-1] == ' ' {
					bytes[i-1] = bytes[i]
					bytes[i] = ' '
				}
			}
		}
	}
	return bytes
}

// ProcessQuote function handles single quotes in the text and adjusts their positions
func ProcessQuote(index int, splitTxt []string, quoteFlag int) int {
	if splitTxt[index] == "'" && index < len(splitTxt)-1 && quoteFlag == 0 {
		// Move single quote to the beginning of the following word
		splitTxt[index+1] = "'" + splitTxt[index+1]
		splitTxt[index] = ""
		quoteFlag = 1
	}

	if splitTxt[index] == "'" && quoteFlag == 1 {
		// Move single quote to the end of the preceding word
		splitTxt[index-1] = splitTxt[index-1] + "'"
		splitTxt[index] = ""
		quoteFlag = 0
	}

	return quoteFlag
}
