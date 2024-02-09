package main

import "github.com/01-edu/z01"

func main() {
	countdown := "9876543210"
	for _, ch := range countdown {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
