package piscine

import "github.com/01-edu/z01"

func QuadC(x,y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 {
					if j == 1 || j == x {
						z01.PrintRune('A')
					} else {
						z01.PrintRune('B')
					}
				} else if i == y {
					if j == 1 || j == x {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('B')
					}
				} else {
					if j == 1 || j == x {
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}
