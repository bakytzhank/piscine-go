package piscine

import (
	"github.com/01-edu/z01"
	"strconv"
)

func ReduceInt(a []int, f func(int, int) int) {
	
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	res := strconv.Itoa(result)
	for _, ch := range res {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
