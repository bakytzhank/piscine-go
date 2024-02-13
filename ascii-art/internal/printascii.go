package internal

import "fmt"

func PrintAscii(input []string) {
	for i:=0; i<len(input); i++ {
		fmt.Println(input[i])
	}
}