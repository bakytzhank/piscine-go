package piscine

func BasicAtoi2(s string) int {
	n := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			n = 0
			break
		} else {
			n = n*10 + int(i-'0')
		}
	}
	return n
}
