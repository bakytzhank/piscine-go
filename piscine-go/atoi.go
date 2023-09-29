package piscine

func Atoi(s string) int {
	n := 0
	sign := 1
	for z, i := range s {
		if z == 0 && i == '+' {
			sign = 1
			continue
		}
		if z == 0 && i == '-' {
			sign = -1
			continue
		}
		if i < '0' || i > '9' {
			n = 0
			break
		} else {
			n = n*10 + int(i-'0')
		}
	}
	return n * sign
}
