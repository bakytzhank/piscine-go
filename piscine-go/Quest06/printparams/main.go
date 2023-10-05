package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	for i := 1; i < len(params); i++ {
		for _, j := range params[i] {
			z01.PrintRune(rune(j))
		}
		z01.PrintRune('\n')
	}
}
