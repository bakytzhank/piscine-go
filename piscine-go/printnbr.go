package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	sign := 1
	if n < 0 {
		sign = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		k := (n / 10) * sign
		if k != 0 {
			PrintNbr(k)
		}
		l := (n % 10 * sign) + '0'
		z01.PrintRune(rune(l))
	} else {
		z01.PrintRune('0')
	}
}
