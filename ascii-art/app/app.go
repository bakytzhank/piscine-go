package app

import (
	"ascii-art/internal"
	"fmt"
	"os"
	"strings"
)

func Run() {
	args := os.Args[1:] // Retrieve command-line arguments
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go $text") // Print usage message if number of arguments are invalid
		return
	}
	internal.GetAsciiArt(strings.ReplaceAll(args[0], `\n`, "\n")) // Generate ASCII art for the provided text replacing 'Enter' from keyboard with "\n"
}
