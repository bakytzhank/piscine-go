package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ConvertToInt(s string) int {
	n := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			n = 0
			break
		} else {
			n = n*10 + int(i-'0')
		}
	}
	return n
}

func main() {
	params := os.Args
	n := len(params)
	arg := 0
	if n == 1 {
		return
	}
	if params[1] == "--upper" {
		for i := 2; i < n; i++ {
			arg = ConvertToInt(params[i])
			if arg >= 1 && arg <= 26 {
				z01.PrintRune(rune(arg + 64))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	} else {
		for i := 1; i < n; i++ {
			arg = ConvertToInt(params[i])
			if arg >= 1 && arg <= 26 {
				z01.PrintRune(rune(arg + 96))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
