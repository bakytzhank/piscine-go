package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]
	if len(params) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(params) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	file, _ := os.Open(params[0])
	text := make([]byte, 14)

	file.Read(text)

	fmt.Println(string(text))

	file.Close()
}
