package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	if len(params) != 1 {
		z01.PrintRune('\n')
		return
	}
	result := ""
	for _, ch := range params[0] {
		if 'a' <= ch && ch <= 'z' {
			result += string('a'+'z'-ch)
		} else if 'A' <= ch && ch <= 'Z' {
			result += string('A'+'Z'-ch)
		} else {
			result += string(ch)
		}
	}

	result += "\n"
	
	for _, ch := range result {
		z01.PrintRune(ch)
	}
}
