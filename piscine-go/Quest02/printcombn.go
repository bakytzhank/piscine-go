package piscine

import "github.com/01-edu/z01"

func Comma(n int, x [9]int, y [9]int) {
	a := 0
	for a < n {
		z01.PrintRune(rune(x[a]) + '0')
		a++
	}
	if x[0] != y[0] {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func Comb() {
	x := [9]int{0}
	y := [9]int{9}
	for x[0] <= y[0] {
		Comma(1, x, y)
		x[0]++
	}
}

func PrintCombN(n int) {
	x := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	y := [9]int{}

	if n == 1 {
		Comb()
	} else {
		a := n - 1
		b := 9
		for a >= 0 {
			y[a] = b
			b--
			a--
		}
		a = n - 1
		for x[0] < y[0] {
			if x[a] != y[a] {
				Comma(n, x, y)
				x[a]++
			}
			if x[a] == y[a] {
				if x[a-1] != y[a-1] {
					Comma(n, x, y)
					x[a-1]++
					b = a
					for b < n {
						x[b] = x[b-1] + 1
						b++
					}
					a = n - 1
				}
			}
			for x[a] == y[a] && x[a-1] == y[a-1] && a > 1 {
				a--
			}
		}
		Comma(n, x, y)
	}
	z01.PrintRune('\n')
}
