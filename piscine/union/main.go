package main

import "os"

func main() {
	params := os.Args[1:]
	if len(params) != 2 {
		os.Stdout.WriteString("\n")
		return
	} 
	input := params[0]+params[1]
	result := ""
	for _, ch1 := range input {
		if result == "" {
			result += string(ch1)
		}
		unique := true
		for _, ch2 := range result {
			if ch1 == ch2 {
				unique = false
				break
			}
		}
		if unique {
			result += string(ch1)
		}
	}
	os.Stdout.WriteString(result+"\n")
}