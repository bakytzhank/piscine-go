package main

import (
	"os"

	"github.com/01-edu/z01"
)

func searchReplace(str, search, replace string) string {
	result := ""
	searchLen := len(search)
	strLen := len(str)

	for i := 0; i < strLen; i++ {
		if i <= strLen-searchLen && str[i:i+searchLen] == search {
			result += replace
			i += searchLen - 1
		} else {
			result += string(str[i])
		}
	}

	return result
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	input := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]

	result := searchReplace(input, search, replace)

	for _, char := range result {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
