package main

import "os"

func main() {
	params := os.Args[1:]
	if len(params) != 1 {
		return
	}
	result := ""
	for _, ch := range params[0] {
		offset := rune(13)
		if 'a' <= ch && ch <='z' || 'A' <= ch && ch <='Z' {
			if 'a' <= ch && ch <='z' && ch + offset > 'z' || 'A' <= ch && ch <='Z' && ch + offset > 'Z'{
				offset = rune(-13)
			}
			result += string(ch+offset)
		} else {
			result +=string(ch)
		}
	}
	os.Stdout.WriteString(result+"\n")
}
