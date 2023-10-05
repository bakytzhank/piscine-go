package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsVowel(c rune) bool {
	return (c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U')
}

func HasTheVowel(r []rune) bool {
	for i := 0; i < len(r); i++ {
		if IsVowel(r[i]) {
			return true
		}
	}
	return false
}

func ReverseVowel(r []rune) []rune {
	j := 0
	vowels := []rune{}
	for i := 0; i < len(r); i++ {
		if IsVowel(r[i]) {
			j++
			vowels = append(vowels, r[i])
		}
	}
	for i := 0; i < len(r); i++ {
		if IsVowel(r[i]) {
			j--
			r[i] = vowels[j]
		}
	}
	return r
}

func main() {
	params := os.Args
	n := len(params)

	if n < 2 {
		z01.PrintRune('\n')
		return
	}
	newString := ""
	for i := 1; i < n; i++ {
		newString += params[i]
		if i < n-1 {
			newString += " "
		}
	}
	Runes := []rune(newString)
	if HasTheVowel(Runes) {
		Runes = ReverseVowel(Runes)
	}
	for _, l := range Runes {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
