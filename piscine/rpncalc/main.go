package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rpncalc(args string) (int, error) {
	stack := []int{}
	operators := strings.Split(args, " ")
	for _, op := range operators {
		if val, err := strconv.Atoi(op); err == nil {
			stack = append(stack, val)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("Error")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch op {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			case "%":
				res = a % b
			default:
				return 0, fmt.Errorf("Error")
			}
			stack = append(stack, res)
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("Error")
	}
	return stack[0], nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	result, err := rpncalc(os.Args[1])
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(result)
}
