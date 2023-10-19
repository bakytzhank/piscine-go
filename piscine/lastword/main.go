package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	if len(params) != 1 {
		return
	}
	word := ""
	last := ""
	for _, ch := range params[0] {
		if ch == ' ' && len(word) > 0 {
			last = word
			word = ""
		} else if ch != ' ' {
			word += string(ch)
		}
	}
	if len(word) > 0 {
		last = word
	}
	if last == "" {
		return
	}
	for _, ch := range last {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
