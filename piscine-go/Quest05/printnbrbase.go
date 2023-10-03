package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	n := len(base)
	for _, l := range base {
		if l == '-' || l == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
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
		nbr = -nbr
	}
	res := ""
	for nbr > 0 {
		res = res + string(base[nbr%n])
		nbr /= n
	}
	for k := len(res) - 1; k >= 0; k-- {
		z01.PrintRune(rune(res[k]))
	}
}
