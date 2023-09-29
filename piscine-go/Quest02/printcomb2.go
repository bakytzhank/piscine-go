package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			l := j + 1
			for k := i; k <= '9'; k++ {
				for ; l <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(32)
					z01.PrintRune(k)
					z01.PrintRune(l)
					if i < '9' || j < '8' || k < '9' || l < '9' {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
				l = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
