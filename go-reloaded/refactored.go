package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filePermission = 0666

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) >= 4 {
		fmt.Println("Too many arguments")
		return
	}

	err := processFile(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
}

func processFile(inputFileName, outputFileName string) error {
	file, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		return err
	}

	txt := string(data)
	txtslice := strings.customSplit(txt)
	printout := ""
	quotechecker := 0

	for index, value := range txtslice {
		processTriggers(index, value, txtslice)

		processPunctuation(index, txtslice)
	}

	printout = strings.Join(txtslice, " ")
	printout = strings.Join(customSplit(printout), " ")

	err = os.WriteFile(outputFileName, []byte(printout), filePermission)
	if err != nil {
		return err
	}

	return nil
}

func processTriggers(index int, value string, txtslice []string) {
	switch value {
	case "(up)":
		txtslice[index-1] = strings.ToUpper(txtslice[index-1])
		txtslice[index] = ""
	case "(up,":
		processCase(index, txtslice, strings.ToUpper)
	case "(low)":
		txtslice[index-1] = strings.ToLower(txtslice[index-1])
		txtslice[index] = ""
	case "(low,":
		processCase(index, txtslice, strings.ToLower)
	case "(cap)":
		txtslice[index-1] = strings.Title(strings.ToLower(txtslice[index-1]))
		txtslice[index] = ""
	case "(cap,":
		processCase(index, txtslice, func(s string) string {
			return strings.Title(strings.ToLower(s))
		})
	case "(hex)":
		processNumericConversion(index, txtslice, 16)
	case "(bin)":
		processNumericConversion(index, txtslice, 2)
	}
}

func processCase(index int, txtslice []string, converter func(string) string) {
	re := regexp.MustCompile("[0-9]+")
	factor := re.FindAllString(txtslice[index+1], -1)
	stringfactor := strings.Join(factor, "")
	intfactor, _ := strconv.Atoi(stringfactor)
	for i := 1; i <= intfactor; i++ {
		txtslice[index-i] = converter(txtslice[index-i])
	}
	txtslice[index+1] = ""
	txtslice[index] = ""
}

func processNumericConversion(index int, txtslice []string, base int) {
	decimal, _ := strconv.ParseInt(txtslice[index-1], base, 32)
	decimalstring := strconv.FormatInt(decimal, 10)
	txtslice[index-1] = decimalstring
	txtslice[index] = ""
}

func processPunctuation(index int, txtslice []string) {
	if txtslice[index] == "'" {
		processQuote(index, txtslice)
	}

	if strings.ContainsAny(txtslice[index], ".!,?;:") {
		if index > 0 && txtslice[index-1] != "" {
			txtslice[index-1] = txtslice[index-1] + txtslice[index]
			txtslice[index] = ""
		}
	}
}

func processQuote(index int, txtslice []string, quotechecker int) {
	if txtslice[index] == "'" && index < len(txtslice)-1 && quotechecker == 0 {
		txtslice[index+1] = "'" + txtslice[index+1]
		txtslice[index] = ""
		quotechecker = 1
		
	}
	if txtslice[index] == "'" && quotechecker == 1   {
	
		txtslice[index-1] = txtslice[index-1] +  "'"
		txtslice[index] = ""
		quotechecker = 0
	}
}

func customSplit(s string) []string {
	var fields []string
	var currentField string

	for _, char := range s {
		if char == ' ' || char == '\t' {
			// Found a whitespace character, add the current field to the slice
			if currentField != "" {
				fields = append(fields, currentField)
				currentField = ""
			}
		} else {
			// Non-whitespace character, add it to the current field
			currentField += string(char)
		}
	}

	// Add the last field (if any)
	if currentField != "" {
		fields = append(fields, currentField)
	}

	return fields
}

