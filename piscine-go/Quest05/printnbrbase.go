package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	n := len(base)
	minusplus := false
	offset := false
	for _, char := range base {
		if char == '+' || char == '-' {
			minusplus = true
		}
	}
	if n < 2 || minusplus {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			nbr = 9223372036854775807
			offset = true
		} else {
			nbr = -nbr
		}
	}
	var res []rune
	for nbr > 0 {
		remainder := nbr % n
		res = append(res, rune(base[remainder]))
		nbr /= n
	}
	if offset {
		res[0]++
	}
	for k := len(res) - 1; k >= 0; k-- {
		z01.PrintRune(res[k])
	}
}
