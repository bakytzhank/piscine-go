package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	argCount := len(args)

	// Print the number of arguments followed by a newline
	z01.PrintRune(rune(argCount) + '0')
	z01.PrintRune('\n')
}