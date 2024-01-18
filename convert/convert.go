package convert

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ProcessCase function processes a trigger with a specified case and applies the conversion function to preceding word
func ProcessCase(index int, splitTxt []string, converter func(string) string) {
	// Check if the field value is not word
	n := 0
	if index != 0 {
		n = 1
	} 
	if (IsWhiteSpaceOrTrigger(splitTxt[index-n]) || IsPunctuation(splitTxt[index-n])) && n < index {
		if n < index {
			n++
		}
		for IsWhiteSpaceOrTrigger(splitTxt[index-n]) || IsPunctuation(splitTxt[index-n]) {
			if n < index {
				n++
			}
		}
	}
	// Apply the conversion function to the preceding words
	splitTxt[index-n] = converter(splitTxt[index-n])

	// Remove the trigger from the processed text
	splitTxt[index] = ""

}

// ProcessCaseN function processes a trigger with a specified case and applies the conversion function to preceding number of words
func ProcessCaseN(index int, splitTxt []string, converter func(string) string) {
	// Extract the numeric value following the trigger
	num := regexp.MustCompile("[0-9]+").FindAllString(splitTxt[index+1], -1)
	stringNum := strings.Join(num, "")
	intNum, _ := strconv.Atoi(stringNum)
	if intNum > index {
		intNum = index
	}
	// Apply the conversion function to the preceding words
	i := 1
	for i <= intNum {
		if (IsWhiteSpaceOrTrigger(splitTxt[index-i]) || IsPunctuation(splitTxt[index-i])) && intNum < index {
			intNum++
		}
		splitTxt[index-i] = converter(splitTxt[index-i])
		i++
	}

	// Remove the trigger and the numeric value from the processed text
	splitTxt[index+1] = ""
	splitTxt[index] = ""
}

// ConvertToDecimal function converts a number from a specified base to decimal
func ConvertToDecimal(index int, splitTxt []string, base int) {
	// Check if the field value is not word
	n := 0
	if index != 0 {
		n = 1
	} 
	if (IsWhiteSpaceOrTrigger(splitTxt[index-n]) || IsPunctuation(splitTxt[index-n])) && n < index {
		if n < index {
			n++
		}
		for IsWhiteSpaceOrTrigger(splitTxt[index-n]) || IsPunctuation(splitTxt[index-n]) {
			if n < index {
				n++
			}
		}
	}
	// Convert the specified number to decimal
	deci, err := strconv.ParseInt(splitTxt[index-n], base, 32)
	deciToStr := strconv.FormatInt(deci, 10)

	// Update the processed text with the decimal value
	if err == nil {
		splitTxt[index-n] = deciToStr
	}
	splitTxt[index] = ""
}

// ConvertAtoAn function converts "a" or "A" to "an" or "An" when followed by a word starting with a vowel or h
func ConvertAtoAn(index int, splitTxt []string) {
	if index < len(splitTxt)-1 {
		// Check if the field value is not word
		n := 1
		if IsWhiteSpaceOrTrigger(splitTxt[index+n]) {
			if n < len(splitTxt)-index-1 {
				n++
			}
			for IsWhiteSpaceOrTrigger(splitTxt[index+n]) {
				if n < len(splitTxt)-index-1 {
					n++
				} else {
					break
				}
			}
		}
		// Check if the next word starts with a vowel or h
		isNextVowel, err := regexp.MatchString(`\A[aeiouh]|\A[AEIOUH]`, splitTxt[index+n])
		if err != nil {
			fmt.Println(err)
			return
		}

		// Replace "a" or "A" with "an" or "An" if the condition is met
		if isNextVowel {
			if splitTxt[index] == "a" {
				splitTxt[index] = "an"
			} else if splitTxt[index] == "A" {
				splitTxt[index] = "An"
			}
		}
	}
}

// Function checks if field is not word
func IsPunctuation (s string) bool {
	return s == "." || s == "," || s == "!" || s == "?" || s == ";" || s == ":" || s == "'"  
}

func IsWhiteSpaceOrTrigger (s string) bool {
	return s == "" || s == "\n" || s == "(cap)" || s == "(up)" || s == "(low)" || s == "(hex)" || s == "(bin)"
}