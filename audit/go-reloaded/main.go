package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	file, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	content := string(file)
	contentwords := SplitWhiteSpaces(content)
	for i, word := range contentwords {
		if word == "(hex)" {
			contentwords[i-1] = hex(contentwords[i-1])
		}
		if word == "(bin)" {
			contentwords[i-1] = bin(contentwords[i-1])
		}
		if strings.HasPrefix(word, "(up") {
			if word[3] == ')' {
				contentwords[i-1] = strings.ToUpper(contentwords[i-1])
			} else if word[3] == ',' {
				numberstr := ""
				for j := 5; word[j] >= '0' && word[j] <= '9'; j++ { //getting the <number> after "(up, "
					numberstr += string(word[j])
				}
				number, err := strconv.Atoi(numberstr)
				if err != nil {

				}
				for k := i - number; k < i; k++ {
					if k < 0 {
						k = 0
					}
					contentwords[k] = strings.ToUpper(contentwords[k])
				}
			}
		}
		if strings.HasPrefix(word, "(low") {
			if word[4] == ')' {
				contentwords[i-1] = strings.ToLower(contentwords[i-1])
			} else if word[4] == ',' {
				numberstr := ""
				for j := 6; word[j] >= '0' && word[j] <= '9'; j++ { //getting the <number> after "(low, "
					numberstr += string(word[j])
				}
				number, err := strconv.Atoi(numberstr)
				if err != nil {

				}
				for k := i - number; k < i; k++ {
					if k < 0 {
						k = 0
					}
					contentwords[k] = strings.ToLower(contentwords[k])
				}
			}
		}
		if strings.HasPrefix(word, "(cap") {
			if word[4] == ')' {
				contentwords[i-1] = Capitalize(contentwords[i-1])
			} else if word[4] == ',' {
				numberstr := ""
				for j := 6; word[j] >= '0' && word[j] <= '9'; j++ { //getting the <number> after "(cap, "
					numberstr += string(word[j])
				}
				number, err := strconv.Atoi(numberstr)
				if err != nil {

				}
				for k := i - number; k < i; k++ {
					if k < 0 {
						k = 0
					}
					contentwords[k] = Capitalize(contentwords[k])
				}
			}
		}
	}

	formatted := ""

	for i, word := range contentwords {
		if i != 0 {
			formatted += " "
		}
		formatted += word
	}

	formatted2 := ""
	for i := 0; i < len(formatted); i++ {
		if i < len(formatted)-1{
			if i != 0 {
				if formatted[i] == ' ' {
					if formatted[i+1] == '.' || formatted[i+1] == ',' || formatted[i+1] == '!' || formatted[i+1] == '?' || formatted[i+1] == ':' || formatted[i+1] == ';' {
						continue
					}
				}
				if (formatted[i-1] == '.' || formatted[i-1] == ',' || formatted[i-1] == '!' || formatted[i-1] == '?' || formatted[i-1] == ':' || formatted[i-1] == ';') && (formatted[i] != '.' && formatted[i] != ',' && formatted[i] != '!' && formatted[i] != '?' && formatted[i] != ':' && formatted[i] != ';' && formatted[i] != ' ') {
					formatted2 += " "
				}
			}
		}
		formatted2 += string(formatted[i])
	}

	formatted2plus := ""
	apostrophes := 0
	for i := 0; i < len(formatted2); i++ {
		if apostrophes%2 == 0 {
			if i < len(formatted2)-1{
				if formatted2[i] == '\'' {
					apostrophes++
					if formatted2[i+1] == ' ' {
						formatted2plus += string(formatted2[i])
						i+=2
					}
				}
			}
		}
		if apostrophes%2 == 1 {
			if i > 0 {
				if formatted2[i] == '\'' {
					apostrophes++
					if formatted2[i-1] == ' ' {
						formatted2plus = formatted2plus[:len(formatted2plus)-1]
						formatted2plus += "'"
						continue
					}
				}
			}
		}
		formatted2plus += string(formatted2[i])
	}


	formatted3 := ""
	openBrackets := 0
	for i := 0; i < len(formatted2plus); i++ {
		if formatted2plus[i] == '(' {
			openBrackets++
		} else if formatted2plus[i] == ')' {
			openBrackets--
		} else {
			if openBrackets == 0 {
				if i == len(formatted2plus)-1 {
					formatted3 += string(formatted2plus[i])
					continue
				}
				if formatted2plus[i+1] == '(' && formatted2plus[i] == ' ' {
					continue
				}
				formatted3 += string(formatted2plus[i])
			}
		}
	}

	formatted3words := SplitWhiteSpaces(formatted3)
	for i, word := range formatted3words {
		if i != len(formatted3words)-1 {
			if word == "a" {
				if formatted3words[i+1][0] == 'a' || formatted3words[i+1][0] == 'e' || formatted3words[i+1][0] == 'i' || formatted3words[i+1][0] == 'o' || formatted3words[i+1][0] == 'u' || formatted3words[i+1][0] == 'h' {
					formatted3words[i] = "an"
					continue
				}
			}
			if word == "A" {
				if formatted3words[i+1][0] == 'a' || formatted3words[i+1][0] == 'e' || formatted3words[i+1][0] == 'i' || formatted3words[i+1][0] == 'o' || formatted3words[i+1][0] == 'u' || formatted3words[i+1][0] == 'h' {
					formatted3words[i] = "An"
					continue
				}
			}
			if word == "an" {
				if formatted3words[i+1][0] != 'a' && formatted3words[i+1][0] != 'e' && formatted3words[i+1][0] != 'i' && formatted3words[i+1][0] != 'o' && formatted3words[i+1][0] != 'u' && formatted3words[i+1][0] != 'h' {
					formatted3words[i] = "a"
				}
			}
			if word == "An" {
				if formatted3words[i+1][0] != 'a' && formatted3words[i+1][0] != 'e' && formatted3words[i+1][0] != 'i' && formatted3words[i+1][0] != 'o' && formatted3words[i+1][0] != 'u' && formatted3words[i+1][0] != 'h' {
					formatted3words[i] = "A"
				}
			}
		}
	}

	formatted3 = ""
	for i, word := range formatted3words {
		if i != 0 {
			formatted3 += " "
		}
		formatted3 += word
	}

	for _, letter := range formatted3 {
		fmt.Print(string(letter))
	}

	err = os.WriteFile(args[1], []byte(formatted3), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}












func Capitalize(s string) string {
	newstring := ""
	counter := 0
	for _, letter := range s {
		if isAlphanumericSpecial(letter) {
			if counter != 0 && letter >= 65 && letter <= 90 {
				newstring += string(letter + 32)
			} else if counter == 0 && letter >= 97 && letter <= 122 {
				newstring += string(letter - 32)
			} else {
				newstring += string(letter)
				counter = 0
			}
			counter++
		} else {
			newstring += string(letter)
			counter = 0
		}
	}
	return newstring
}

func isAlphanumericSpecial(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}

func SplitWhiteSpaces(s string) []string {
	var slice []string
	word := ""
	openedBrackets := 0
	for i, letter := range s {
		if letter == '(' {
			openedBrackets++
		}
		if letter == ')' {
			openedBrackets--
		}
		if (letter == ' ' || letter == '\n' || i == len(s)-1) && openedBrackets == 0 {
			if i == len(s)-1 && letter != ' ' {
				word += string(letter)
			}
			if word != "" {
				slice = append(slice, word)
			}
			word = ""
			continue
		}
		word += string(letter)
	}
	return slice
}

func bin(s string) string {
	result := 0
	for i, letter := range s {
		if letter == '0' || letter == '1' {
			bit, err := strconv.Atoi(string(letter))
			if err != nil {
				return "0"
			}
			result += bit * int(math.Pow(2, float64(len(s)-i-1)))
		} else {
			return "0"
		}
	}
	return strconv.Itoa(result)
}

func hex(s string) string {
	result := 0
	for i, letter := range s {
		if (letter >= '0' && letter <= '9') || (letter >= 'A' && letter <= 'F') || (letter >= 'a' && letter <= 'f') {
			bit, err := strconv.Atoi(string(letter))
			if err != nil {

			}
			if letter >= 'A' && letter <= 'F' {
				bit = int(letter - 55) //so that A is 10, B is 11 etc.
			}
			if letter >= 'a' && letter <= 'f' {
				bit = int(letter - 87) // so that a is 10, b is 11 etc.
			}
			result += bit * int(math.Pow(16, float64(len(s)-i-1)))
		} else {
			return "0"
		}
	}
	return strconv.Itoa(result)
}
