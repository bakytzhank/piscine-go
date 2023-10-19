package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	if len(params) != 2 {
		return
	}
	first := params[0]
	second := params[1]
	result := ""
	contains := false
	for _, ch1 := range first {
		for _, ch2 := range second {
			if ch1 == ch2 {
				contains = false
				for _, ch := range result {
					if ch1 == ch {
						contains = true
						break
					}
				} 
				if !contains {
					result += string(ch1)
				}
			}
		}
	}
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
