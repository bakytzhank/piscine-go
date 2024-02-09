package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    input := os.Args[1:]
    if len(input) != 1 {
        return
    }
    if input[0] == "" {
        z01.PrintRune('\n')
        return
    }
    var result []string
    word := ""
    for _, ch := range input[0] {
        if ch == ' ' && len(word)>0 {
            result = append(result, word)
            word = ""
        } else if ch != ' ' {
            word += string(ch)
        }
    }
    if len(word) > 0 {
        result = append(result, word)
    }
    for i:=len(result)-1; i > 0; i-- {
        for _, ch := range result[i] {
            z01.PrintRune(ch)
        }
        z01.PrintRune(' ')        
    }
    for _, ch := range result[0] {
        z01.PrintRune(ch)
    }
    z01.PrintRune('\n')
}
