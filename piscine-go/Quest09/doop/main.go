package main

import (
	"os"
)

func atoi(s string) (int, bool) {
	n := 0
	sign := 1
	for _, char := range s {
		if char == '-' {
			sign = -1
		} else if char >= '0' && char <= '9' {
			digit := int(char - '0')
			if n > (1<<63-1-digit)/10 {
				// Overflow would occur
				return 0, false
			}
			n = n*10 + digit
		} else {
			return 0, false
		}
	}
	return n * sign, true
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}
	return sign + result
}

func safeAdd(a, b int) (int, bool) {
	if (b > 0 && a > (1<<62)-b) || (b < 0 && a < -(1<<62)-b) {
		return 0, true
	}
	return a + b, false
}

func safeSubtract(a, b int) (int, bool) {
	if (b > 0 && a < -(1<<62)+b) || (b < 0 && a > (1<<62)+b) {
		return 0, true
	}
	return a - b, false
}

func safeMultiply(a, b int) (int, bool) {
	if a > 0 {
		if b > 0 && a > (1<<(63/2))/b {
			return 0, true
		}
		if b < 0 && b < -(1<<(63/2))/a {
			return 0, true
		}
	}
	if a < 0 {
		if b > 0 && a < -(1<<(63/2))/b {
			return 0, true
		}
		if b < 0 && b < (1<<(63/2))/a {
			return 0, true
		}
	}
	return a * b, false
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	a, validA := atoi(args[0])
	op := args[1]
	b, validB := atoi(args[2])

	if !validA || !validB {
		return
	}

	var result int
	overflow := false

	switch op {
	case "+":
		result, overflow = safeAdd(a, b)
	case "-":
		result, overflow = safeSubtract(a, b)
	case "*":
		result, overflow = safeMultiply(a, b)
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = a % b
	default:
		return
	}

	if overflow {
		return
	}

	resultStr := itoa(result)
	os.Stdout.WriteString(resultStr + "\n")
}
