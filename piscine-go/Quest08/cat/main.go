package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	executablePath := os.Args[0]
	executableDir := strings.TrimSuffix(executablePath, "/cat/main")

	if len(os.Args) == 1 {
		// If no arguments are provided, read from stdin and print to stdout.
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
	} else {
		// If arguments are provided, open and read each file, then print its contents.
		for _, filename := range os.Args[1:] {
			filePath := executableDir + "/" + filename
			file, err := os.Open(filePath)
			if err != nil {
				os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			_, err = io.Copy(os.Stdout, file)
			if err != nil {
				os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			file.Close()
		}
	}
}
