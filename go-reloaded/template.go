package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := processLine(line)
		outputFile.WriteString(modifiedLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println("Processing completed. Check", outputFileName, "for the result.")
}

func processLine(line string) string {
	words := strings.Fields(line)

	for i := 0; i < len(words); i++ {
		word := words[i]

		switch word {
		case "(hex)":
			if i > 0 {
				decimal, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
				}
			}
		case "(bin)":
			if i > 0 {
				decimal, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(decimal, 10)
				}
			}
		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = strings.Title(words[i-1])
			}
		case "(low,":
			if i+1 < len(words) && isNumber(words[i+1]) {
				count, _ := strconv.Atoi(words[i+1])
				i++
				for j := 0; j < count && i < len(words); j++ {
					words[i] = strings.ToLower(words[i])
					i++
				}
				i--
			}
		case "(up,":
			if i+1 < len(words) && isNumber(words[i+1]) {
				count, _ := strconv.Atoi(words[i+1])
				i++
				for j := 0; j < count && i < len(words); j++ {
					words[i] = strings.ToUpper(words[i])
					i++
				}
				i--
			}
		case "(cap,":
			if i+1 < len(words) && isNumber(words[i+1]) {
				count, _ := strconv.Atoi(words[i+1])
				i++
				for j := 0; j < count && i < len(words); j++ {
					words[i] = strings.Title(words[i])
					i++
				}
				i--
			}
		default:
			if strings.HasPrefix(word, "'") && strings.HasSuffix(word, "'") {
				words[i] = strings.Trim(word, "'")
			}
			if len(word) == 1 && word[0] == 'a' && i+1 < len(words) && isVowelOrH(words[i+1]) {
				words[i] = "an"
			}
		}
	}

	return formatPunctuation(strings.Join(words, " "))
}

func formatPunctuation(text string) string {
	text = regexp.MustCompile(`(\w)([.,!?;:])`).ReplaceAllString(text, "$1 $2")
	text = regexp.MustCompile(`\.{3}`).ReplaceAllString(text, "...")
	text = regexp.MustCompile(`([!?])\.`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`([!?])\?`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`'\s*(\w+)\s*'`).ReplaceAllString(text, "'$1'")
	return text
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isVowelOrH(s string) bool {
	if len(s) > 0 {
		firstChar := unicode.ToLower(rune(s[0]))
		return firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'h'
	}
	return false
}