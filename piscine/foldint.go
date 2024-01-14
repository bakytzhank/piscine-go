package piscine

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	res := n 
	for _, i := range a {
		res = f(res, i)
	}
	fmt.Println(res)
}