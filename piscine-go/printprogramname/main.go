package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	count := 0
	for _, l := range name {
		if count > 1 {
			z01.PrintRune(rune(l))
		}
		count++
	}
	z01.PrintRune('\n')
}
