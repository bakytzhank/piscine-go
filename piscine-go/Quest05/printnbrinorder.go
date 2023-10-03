package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := [10]int{}

	num := n
	for num > 0 {
		digit := num % 10
		digits[digit]++
		num /= 10
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}
	}
}
