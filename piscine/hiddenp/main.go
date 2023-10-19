package main

import (
	"os"
	"github.com/01-edu/z01"
)

func hidden(s1, s2 string) bool {
	indexS2 := 0
	for _, ch1 := range s1 {
		found := false
		for indexS2 < len(s2) {
			ch2 := rune(s2[indexS2])
			indexS2++
			if ch1 == ch2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	params := os.Args[1:]
	if len(params) != 2 {
		return
	}
	s1 := params[0]
	s2 := params[1]
	
	if hidden(s1, s2) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}