package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	if len(params) != 1 || len(params[0]) == 0 {
		return
	}
	input := params[0]
	result := ""
	vowels := []rune{'a','e','i','o','u','A','E','I','O','U'}
	for _, r := range vowels {
		if r == rune(input[0]) {
			result = input + "ay\n"
		}
	}
	if result == "" {
		n := 0
		for i, ch := range input {
			for _, r := range vowels {
				if ch == r {
					n = i
					break
				}
			}
			if n != 0 {
				break
			}
		}
		if n != 0 {
			result = input[n:] + input[:n] + "ay\n"
		} else {
			result = "No vowels\n"
		}
	}
	for _, ch := range result {
		z01.PrintRune(ch)
	}
}
