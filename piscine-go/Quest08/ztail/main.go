package main

import (
	"fmt"
	"os"
)

func parseNumChars(numChars string) int {
	n := 0
	for _, digit := range numChars {
		if digit >= '0' && digit <= '9' {
			n = n*10 + int(digit-'0')
		} else {
			break
		}
	}
	return n
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func processFile(filename, numChars string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	fileSize := fileStat.Size()
	num := parseNumChars(numChars)

	if fileSize < int64(num) {
		num = int(fileSize)
	}

	buffer := make([]byte, num)
	_, err = file.ReadAt(buffer, fileSize-int64(num))
	if err != nil {
		return err
	}

	fmt.Print(string(buffer))
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . -c <num> file1.txt [file2.txt ...]")
		os.Exit(1)
	}

	numChars := os.Args[2]
	filenames := os.Args[3:]

	exitStatus := 0

	for i, filename := range filenames {
		if i == 0 && fileExists(filename) {
			fmt.Printf("==> %s <==\n", filename)
			if err := processFile(filename, numChars); err != nil {
				fmt.Println(err)
				exitStatus = 1
			}
		} else if i != 0 && fileExists(filename) {
			fmt.Println()
			fmt.Printf("==> %s <==\n", filename)
			if err := processFile(filename, numChars); err != nil {
				fmt.Println(err)
				exitStatus = 1
			}
		} else if err := processFile(filename, numChars); err != nil {
			fmt.Fprintln(os.Stderr, err)
			exitStatus = 1
		}
	}

	os.Exit(exitStatus)
}
