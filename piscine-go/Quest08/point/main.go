package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PointsToString(num int) string {
	var str []rune
	for num > 0 {
		digit := num % 10
		char := rune('0' + digit)
		str = append([]rune{char}, str...)
		num /= 10
	}
	return string(str)
}

func main() {
	points := point{}

	setPoint(&points)

	outputStr := "x = " + PointsToString(points.x) + ", y = " + PointsToString(points.y) + "\n"
	for _, ch := range outputStr {
		z01.PrintRune(ch)
	}
}
