package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestProcessText(t *testing.T) {
	for i := 1; i <= 20; i++ {
		inputFile := fmt.Sprintf("testcase%d.txt", i)
		expectedFile := fmt.Sprintf("expected%d.txt", i)

		t.Run(inputFile, func(t *testing.T) {
			// Prepare a temporary file for testing
			tempFile, err := ioutil.TempFile("", "testfile.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Read test data from the input file
			testData, err := ioutil.ReadFile(inputFile)
			if err != nil {
				t.Fatal(err)
			}

			// Write test data to the temporary file
			err = ioutil.WriteFile(tempFile.Name(), testData, 0666)
			if err != nil {
				t.Fatal(err)
			}

			// Call the main function with the temporary file and output file arguments
			os.Args = []string{"cmd", tempFile.Name(), "output.txt"}
			main()

			// Read the modified text from the output file
			modifiedData, err := ioutil.ReadFile("output.txt")
			if err != nil {
				t.Fatal(err)
			}

			// Read the expected result from the test file
			expectedResult, err := ioutil.ReadFile(expectedFile)
			if err != nil {
				t.Fatal(err)
			}

			// Compare the actual result with the expected result
			if string(modifiedData) != string(expectedResult) {
				t.Errorf("Test case %s failed. Expected:\n%s\nGot:\n%s", inputFile, string(expectedResult), string(modifiedData))
			}
		})
	}
}

// Add more tests as needed for different scenarios and edge cases
