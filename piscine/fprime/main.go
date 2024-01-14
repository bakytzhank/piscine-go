// package main

// import (
// 	"os"
// 	"strconv"
// 	"fmt"
// )

// func main() {
// 	params := os.Args[1:]
// 	if len(params) != 1 {
// 		return
// 	}
// 	n, err := strconv.Atoi(params[0])
// 	if n <= 1 || err != nil {
// 		fmt.Println()
// 		return
// 	}
// 	factors := []int{}
// 	divisor := 2
// 	for n > 1 {
// 		if n%divisor == 0 {
// 			factors := append(factors, divisor)
// 			n = n/divisor
// 		} else {
// 			divisor++
// 		}
// 	}
// 	for i := 0; i < len(factors)-1; i++ {
// 		fmt.Print(factors[i],"*")
// 	}
// 	fmt.Print(factors[len(factors)-1],"\n")
// }

package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Fprime(n int) {
	if n <= 1 {
		return
	}

	for i := 2; i <= n; i++ {
		if isPrime(i) && n%i == 0 {
			for n%i == 0 {
				n /= i
				fmt.Printf("%d", i)
				if n > 1 {
					fmt.Printf("*")
				}
			}
		}
	}
	fmt.Printf("\n")
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	Fprime(num)
}