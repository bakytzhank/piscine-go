package main

import (
	"os"
	"fmt"
	"strconv"
)

func main () {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println(a)
}
