package main

import (
	"os"
	"fmt"
)

func isValidOption(s string) bool {
	if len(s) < 2 || s[0] != '-' {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}

func displayOptions() {
	os.Stdout.WriteString("options: abcdefghijklmnopqrstuvwxyz\n")
}

func parseOptions(args []string) {
	var options int32

	for _, arg := range args {
		if !isValidOption(arg) {
			os.Stdout.WriteString("Invalid Option\n")
			return
		}
		if arg == "-h" || arg[0] =='-' && arg[1] == 'h' {
			displayOptions()
			return
		}
		for i := 1; i < len(arg); i++ {
			pos := int(arg[i]-'a')
			options |= 1 << pos
		}
	}
	fmt.Printf("%08b %08b %08b %08b\n", byte(options>>24), byte(options>>16), byte(options>>8), byte(options))
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		displayOptions()
		return
	}
	parseOptions(args)
}