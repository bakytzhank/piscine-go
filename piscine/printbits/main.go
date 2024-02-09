package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func printBits(num int) {
	for i := 7; i >= 0; i-- {
		bit := (num >> uint(i)) & 1
		z01.PrintRune(rune('0' + bit))
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		for i := 0; i < 8; i++ {
			z01.PrintRune('0')
		}
		return
	}

	printBits(num)
}
