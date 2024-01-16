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
	n := 1
	if IsNotWord(splitTxt[index-1]) {
		n = 2
		for IsNotWord(splitTxt[index-n]) {
			n++
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

	// Apply the conversion function to the preceding words
	i := 1
	for i <= intNum {
		if IsNotWord(splitTxt[index-i]) {
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
	n := 1
	if IsNotWord(splitTxt[index-1]) {
		n = 2
		for IsNotWord(splitTxt[index-n]) {
			n++
		}
	}
	// Convert the specified number to decimal
	deci, _ := strconv.ParseInt(splitTxt[index-n], base, 32)
	deciToStr := strconv.FormatInt(deci, 10)

	// Update the processed text with the decimal value
	splitTxt[index-n] = deciToStr
	splitTxt[index] = ""
}

// ConvertAtoAn function converts "a" or "A" to "an" or "An" when followed by a word starting with a vowel or h
func ConvertAtoAn(index int, splitTxt []string) {
	if index < len(splitTxt)-1 {
		// Check if the next word starts with a vowel or h
		isNextVowel, err := regexp.MatchString(`\A[aeiouh]|\A[AEIOUH]`, splitTxt[index+1])
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

func IsNotWord (s string) bool {
	return s == "" || s == "\n" || s == "." || s == "," || s == "!" || s == "?" || s == ";" || s == ":" || s == "'"
}
