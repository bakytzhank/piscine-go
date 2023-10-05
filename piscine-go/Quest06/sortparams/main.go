package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	n := len(params)
	for i := 1; i < n-1; i++ {
		swapped := false
		for j := 1; j < n-i; j++ {
			if rune(params[j][0]) > rune(params[j+1][0]) {
				params[j], params[j+1] = params[j+1], params[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	for i := 1; i < len(params); i++ {
		for _, j := range params[i] {
			z01.PrintRune(rune(j))
		}
		z01.PrintRune('\n')
	}
}
