package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// If no arguments are provided, read from stdin and write to stdout
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
	} else {
		// If arguments are provided, treat them as file paths
		for _, filePath := range os.Args[1:] {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				continue
			}
			defer file.Close()

			_, err = io.Copy(os.Stdout, file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			}
		}
	}
}
