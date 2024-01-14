package main

import (
    "fmt"
    "os"
)

func atoi(s string) int {
    num := 0
    for index := range s {
        if s[index] < '0' || s[index] > '9' {
            return 0
        } else {
            num = num*10 + int(s[index]-'0')
        }
    }
    return num
}

func main() {
    if len(os.Args) == 2 && atoi(os.Args[1]) > 0 && atoi(os.Args[1]) < 4000 {
        arg := os.Args[1]
        ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
        tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
        hrns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
        ths := []string{"", "M", "MM", "MMM"}
        x := []string{ths[atoi(arg)/1000], hrns[(atoi(arg)%1000)/100], tens[(atoi(arg)%100)/10], ones[atoi(arg)%10]}
        cal := []rune{}
        for _, str := range x {
            if str != "IX" && str != "XC" && str != "CM" && str != "IV" && str != "XL" && str != "CD" {
                for _, v := range str {
                    cal = append(cal, v, '+')
                }
            } else {
                cal = append(cal, '(')
                r := []rune(str)
                for i, j := 0, len(r)-1; j >= len(r)/2; i, j = i+1, j-1 {
                    r[i], r[j] = r[j], r[i]
                }
                for i, v := range r {
                    cal = append(cal, v)
                    if i != len(str)-1 {
                        cal = append(cal, '-')
                    } else {
                        cal = append(cal, ')')
                    }
                }
                cal = append(cal, '+')

            }
        }
        fmt.Printf("%s \n", string(cal[:len(cal)-1]))
        for _, v := range x {
            fmt.Printf(v)
        }
        fmt.Printf("\n")
    } else {
        fmt.Printf("ERROR: cannot convert to roman digit\n")
    }
}