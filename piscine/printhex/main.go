package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	if len(params) != 1 {
		return
	}
	E := "ERROR\n"
	dec, error := strconv.Atoi(params[0])
	if error != nil || dec < 0 {
		for _, ch := range E {
			z01.PrintRune(ch)
		}
		return
	}
	base := "0123456789abcdef"
	result := ""
	if dec == 0 {
		result = "0"
	}
	for dec > 0 {
		digit := dec%16
		result = string(base[digit]) + result
		dec /= 16
	}
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}