package main

import (
	"fmt"
	"os"
	"go-reloaded/processors"
	"go-reloaded/strprocess"
)

const filePermission = 0666 // File Permission set to Read & Write

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) < 3 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	}

	// Process the input file and write the result to the output file
	err := processFile(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
}

// processFile function reads from the input file, processes the content, and writes the result to the output file
func processFile(inputFileName, outputFileName string) error {
	// Open the input file
	file, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the content of the input file
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		return err
	}

	// Split the text into an array of words
	splitTxt := strprocess.CustomSplit(string(data))

	// Process various triggers in the text
	resultTxt := ""
	quoteFlag := 0
	for index, val := range splitTxt {
		processors.ProcessTriggers(index, val, splitTxt)
		quoteFlag = processors.ProcessQuote(index, splitTxt, quoteFlag)
	}

	// Handle punctuation and join the processed text
	resultTxt = strprocess.CustomJoin(strprocess.CustomSplit(string(processors.ProcessPunctuation(splitTxt))))

	// Write the result to the output file
	err = os.WriteFile(outputFileName, []byte(resultTxt), filePermission)
	if err != nil {
		return err
	}

	return nil
}