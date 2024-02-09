package main

import (
	"github.com/01-edu/z01"
	"unicode"
)

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		if i > 0 {
			z01.PrintRune(' ')
		}
		hex := "0123456789abcdef"
		z01.PrintRune(rune(hex[b/16]))
		z01.PrintRune(rune(hex[b%16]))
	}
	z01.PrintRune('\n')
	for _, b := range arr {
		//if b >= 32 && b <= 126 {
		if unicode.IsGraphic(rune(b)) {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}