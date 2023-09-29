package piscine

func BasicAtoi(s string) int {
	n := 0
	for _, i := range s {
		n = n*10 + int(i-'0')
	}
	return n
}
