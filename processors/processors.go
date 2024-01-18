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
		convert.ProcessCase(index, splitTxt, strings.ToUpper)
	case "(up,":
		convert.ProcessCaseN(index, splitTxt, strings.ToUpper)
	case "(low)":
		convert.ProcessCase(index, splitTxt, strings.ToLower)
	case "(low,":
		convert.ProcessCaseN(index, splitTxt, strings.ToLower)
	case "(cap)":
		convert.ProcessCase(index, splitTxt, func(s string) string {
			return strings.Title(strings.ToLower(s))
		})
	case "(cap,":
		convert.ProcessCaseN(index, splitTxt, func(s string) string {
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
					if i != len(bytes)-1 {
						if bytes[i+1] == '\'' {
							bytes[i] = bytes[i+1]
							bytes[i+1] = ' '
						}
					}
				}
			}
		}
	}
	return bytes
}

// ProcessQuote function handles single quotes in the text and adjusts their positions
func ProcessQuote(index int, splitTxt []string, quoteFlag int, nextSkip int) (int, int) {
	if nextSkip == 1 {
		nextSkip = 0
		return quoteFlag, nextSkip
	}
	if quoteFlag == 0 {
		if index < len(splitTxt)-1 {
			// If single quote has not been opened and not adjusted, move it to the beginning of the following word
			if splitTxt[index] == "'" {
				// Move single quote to the beginning of the following word
				splitTxt[index+1] = "'" + splitTxt[index+1]
				splitTxt[index] = ""
				quoteFlag = 1
				nextSkip = 1
			} else if strings.HasPrefix(splitTxt[index], "'") {
				quoteFlag = 1
			} else if strings.HasSuffix(splitTxt[index], "'") {
				splitTxt[index+1] = "'" + splitTxt[index+1]
				splitTxt[index] = strings.TrimSuffix(splitTxt[index], "'")
				quoteFlag = 1
				nextSkip = 1
			}
		}
	} else {
		if index > 0 {
			// If single quote has been opened and not adjusted, move it to the end of the preceding word
			if splitTxt[index] == "'" {
				
				splitTxt[index-1] = splitTxt[index-1] + "'"
				splitTxt[index] = ""
				quoteFlag = 0
			} else if strings.HasSuffix(splitTxt[index], "'") {
				quoteFlag = 0
			} else if strings.HasPrefix(splitTxt[index], "'") {
				splitTxt[index-1] = splitTxt[index-1] + "'"
				splitTxt[index] = strings.TrimPrefix(splitTxt[index], "'")
				quoteFlag = 0
			}
		}
	}

	return quoteFlag, nextSkip
}