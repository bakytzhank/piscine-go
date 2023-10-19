package main

import (
	"os"
	"github.com/01-edu/z01"
)

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	n := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			n = n*10 + int(ch-'0')
		} else {
			return 0
		}
	}
	return n
}

func itoa(n int) string {
	res := ""
	for n > 0 {
		digit := n%10
		res = string(digit+'0') + res
		n /= 10
	}
	return res
}

func main() {
	params := os.Args[1:]
	if len(params) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	n := atoi(params[0])
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	sum := 0
	for i := 2; i <= n; i++ {
		if isprime(i) {
			sum += i
		}
	}
	result := itoa(sum)
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
