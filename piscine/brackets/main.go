package main

import (
	"fmt"
	"os"
)

func isMatchingPair(left rune, right rune) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '{' && right == '}' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	return false
}

func isBalanced(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || !isMatchingPair(stack[len(stack)-1], char) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println()
	} else {
		for _, arg := range args {
			if isBalanced(arg) {
				fmt.Println("OK")
			} else {
				fmt.Println("Error")
			}
		}
	}
}