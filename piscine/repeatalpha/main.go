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
	s := params[0]
	var n int
	for _, ch := range s {
		if 'a' <= ch && ch <= 'z' {
			n = int(ch - 'a' + 1)
			for i := 0; i < n; i++ {
				z01.PrintRune(ch)
			} 
		} else if 'A' <= ch && ch <= 'Z' {
			n = int(ch - 'A' + 1)
			for i := 0; i < n; i++ {
				z01.PrintRune(ch)
			} 
		} else {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}