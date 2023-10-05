package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func HaveInString(s string, toFind string) bool {
	for i := 0; i < len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return true
		}
	}
	return false
}

func Insert(s string, n int) string {
	newString := ""
	for i := n; i < len(s); i++ {
		newString += string(s[i])
	}
	return newString
}

func SortString(s string) string {
	runes := []rune(s)
	n := len(runes)
	newString := ""
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	for i := 0; i < n; i++ {
		newString += string(runes[i])
	}
	return newString
}

func main() {
	params := os.Args
	n := len(params)
	var resultString string
	if n == 1 || params[1] == "-h" || params[1] == "--help" {
		fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
		return
	}
	if HaveInString(params[1], "--insert") {
		if params[2] == "--order" || params[2] == "-o" {
			resultString = SortString(params[3] + Insert(params[1], 9))
		} else {
			resultString = params[2] + Insert(params[1], 9)
		}
	} else if HaveInString(params[1], "-i") {
		if params[2] == "--order" || params[2] == "-o" {
			resultString = SortString(params[3] + Insert(params[1], 3))
		} else {
			resultString = params[2] + Insert(params[1], 3)
		}
	} else if params[1] == "--order" || params[1] == "-o" {
		resultString = SortString(params[2])
	} else {
		resultString = params[1]
	}
	for _, l := range resultString {
		z01.PrintRune(rune(l))
	}
	z01.PrintRune('\n')
}
