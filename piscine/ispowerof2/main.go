package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPowerOf2(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	result := isPowerOf2(num)
	fmt.Println(result)
}