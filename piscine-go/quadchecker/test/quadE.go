package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	n := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			n = 0
			break
		} else {
			n = n*10 + int(i-'0')
		}
	}
	return n
}

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 {
					if j == 1 {
						fmt.Print("A")
						continue
					}
					if j == x {
						fmt.Print("C")
						continue
					}
					fmt.Print("B")
					continue
				}
				if i == y {
					if j == 1 {
						fmt.Print("C")
						continue
					}
					if j == x {
						fmt.Print("A")
						continue
					}
					fmt.Print("B")
					continue
				}
				if j == 1 || j == x {
					fmt.Print("B")
					continue
				}
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
	}
}

func main() {
	args := os.Args[1:]

	x := Atoi(args[0])
	y := Atoi(args[1])

	QuadE(x, y)
}
