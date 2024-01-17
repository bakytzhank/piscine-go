package main

import "os"
import "github.com/01-edu/z01"

func BrainfuckInterpret(code string) {
	var tape [2048]byte
	ptr := 0

	loopStack := []int{}
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			z01.PrintRune(rune(tape[ptr]))
		case '[':
			if tape[ptr] == 0 {
				loopCount := 1
				for loopCount > 0 {
					i++
					if code[i] == ']' {
						loopCount--
					} else if code[i] == '[' {
						loopCount++
					}
				}
			} else {
				loopStack = append(loopStack, i)
			}
		case ']':
			i = loopStack[len(loopStack)-1] - 1
			loopStack = loopStack[:len(loopStack)-1]
		}
	}
}

func main() {
	if len(os.Args) == 2 {
		code := os.Args[1]
		BrainfuckInterpret(code)
	}
}
