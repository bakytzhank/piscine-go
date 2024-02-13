package internal

import (
	"bufio"
	"fmt"
	"os"
	"hash/fnv"
)

func GetBanners() [][]string {
	filepath := "banners/standard.txt"
	initialHash := uint32(957876679) // Initial hash value to detect changes in the file
	
	// Calculate the current hash of the file
	currentHash, err := CalculateHash(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	// Check if the file has been changed since the initial hash
	if currentHash != initialHash {
		fmt.Println("File has been changed")
		return nil
	}
	// Process file and return errors if any
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile) // Create a scanner to read the file
	var result [][]string
	var ch []string

	for fileScanner.Scan() {
		line := fileScanner.Text() // Read a line from the file
		if line == "" {
			if len(ch) > 0 {
				result = append(result, ch) // Add character's ASCII art to the result
				ch = nil
			}
		} else {
			ch = append(ch, line) // Add line to character's ASCII art
		}
	}
	if len(ch) > 0 {
		result = append(result, ch) // Add the last character's ASCII art if not empty
	}

	return result
}

// Calculates the hash value of a file located at the given filepath
func CalculateHash(filepath string) (uint32, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	h := fnv.New32() // Create a new FNV-1a hash
	_, err = h.Write([]byte{})
	if err != nil {
		return 0, err
	}

	// Read the file and update the hash
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && n == 0 {
			break
		} else if err != nil {
			return 0, err
		}
		h.Write(buffer[:n])
	}

	return h.Sum32(), nil
}