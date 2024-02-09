package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <string1> <string2>")
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	result := wdmatch(str1, str2)

	fmt.Println(result)
}

func wdmatch(str1, str2 string) string {
	index1 := 0
	index2 := 0

	for index2 < len(str2) {
		if str1[index1] == str2[index2] {
			index1++
			if index1 == len(str1) {
				return "1"
			}
		}
		index2++
	}

	return "0"
}
