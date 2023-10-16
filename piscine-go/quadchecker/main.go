package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Setting up a scanner to read input from the command line
	reader := bufio.NewScanner(os.Stdin)
	var inputRunes []rune
	inputString := ""
	outputString := []string{}

	// Reading input until there is no more data
	for reader.Scan() {
		inputString = inputString + reader.Text() + "\n"
	}

	// Converting the input string into an array of runes
	for _, ch := range inputString {
		inputRunes = append(inputRunes, ch)
	}

	// Getting the width and height of the input string
	x, y := GetXY(inputRunes)
	X := Itoa(x)
	Y := Itoa(y)

	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Checking the input string against each quad and appending results accordingly
	if QuadA(x, y) == inputString {
		outputString = append(outputString, "[quadA] [" + X + "] [" + Y + "]")
	}
	if QuadB(x, y) == inputString {
		outputString = append(outputString, "[quadB] [" + X + "] [" + Y + "]")
	}
	if QuadC(x, y) == inputString {
		outputString = append(outputString, "[quadC] [" + X + "] [" + Y + "]")
	}
	if QuadD(x, y) == inputString {
		outputString = append(outputString, "[quadD] [" + X + "] [" + Y + "]")
	}
	if QuadE(x, y) == inputString {
		outputString = append(outputString, "[quadE] [" + X + "] [" + Y + "]")
	}

	// Printing the output or indicating if the input string doesn't match any quad
	if len(outputString) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	for i := 0; i < len(outputString)-1; i++ {
		fmt.Print(outputString[i] + " || ")
	}
	fmt.Println(outputString[len(outputString)-1])
}

// GetXY function calculates the width and height of the input string
func GetXY(output []rune) (x, y int) {
	countX := 0
	countY := 0
	flag := true
	for _, s := range output {
		if s == '\n' {
			countY++
			flag = false
			continue
		}
		if flag == true {
			countX++
		}
	}
	return countX, countY
}

// Itoa converts an integer to its corresponding string representation
func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}
	return result
}

// This function generates the pattern layout for quad A
func QuadA(x, y int) string {
	result := ""
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 || i == y {
					if j == 1 || j == x {
						result += "o"
						continue
					}
					result += "-"
					continue
				}
				if j == 1 || j == x {
					result += "|"
					continue
				}
				result += " "
			}
			result += "\n"
		}
	}
	return result
}

// This function generates the pattern layout for quad B
func QuadB(x, y int) string {
	result := ""
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 {
					if j == 1 {
						result += "/"
						continue
					}
					if j == x {
						result += "\\"
						continue
					}
					result += "*"
					continue
				}
				if i == y {
					if j == 1 {
						result += "\\"
						continue
					}
					if j == x {
						result += "/"
						continue
					}
					result += "*"
					continue
				}
				if j == 1 || j == x {
					result += "*"
					continue
				}
				result += " "
			}
			result += "\n"
		}
	}
	return result
}

// This function generates the pattern layout for quad C
func QuadC(x, y int) string {
	result := ""
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 {
					if j == 1 || j == x {
						result += "A"
						continue
					}
					result += "B"
					continue
				}
				if i == y {
					if j == 1 || j == x {
						result += "C"
						continue
					}
					result += "B"
					continue
				}
				if j == 1 || j == x {
					result += "B"
					continue
				}
				result += " "
			}
			result += "\n"
		}
	}
	return result
}

// This function generates the pattern layout for quad D
func QuadD(x, y int) string {
	result := ""
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 || i == y {
					if j == 1 {
						result += "A"
						continue
					}
					if j == x {
						result += "C"
						continue
					}
					result += "B"
					continue
				}
				if j == 1 || j == x {
					result += "B"
					continue
				}
				result += " "
			}
			result += "\n"
		}
	}
	return result
}

// This function generates the pattern layout for quad E
func QuadE(x, y int) string {
	result := ""
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 {
					if j == 1 {
						result += "A"
						continue
					}
					if j == x {
						result += "C"
						continue
					}
					result += "B"
					continue
				}
				if i == y {
					if j == 1 {
						result += "C"
						continue
					}
					if j == x {
						result += "A"
						continue
					}
					result += "B"
					continue
				}
				if j == 1 || j == x {
					result += "B"
					continue
				}
				result += " "
			}
			result += "\n"
		}
	}
	return result
}
