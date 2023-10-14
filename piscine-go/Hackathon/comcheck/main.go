package main

import "os"

func main() {
	params := os.Args[1:]
	for i := 0; i < len(params); i++ {
		if params[i] == "01" || params[i] == "galaxy" || params[i] == "galaxy 01" {
			os.Stdout.WriteString("Alert!!!\n")
			break
		}
	}
}
