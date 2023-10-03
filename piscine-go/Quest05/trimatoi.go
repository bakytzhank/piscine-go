package piscine

func TrimAtoi(s string) int {
	sign := 1
	n := 0
	digit := false
	for _, l := range s {
		if digit && l == 48 {
			n = n*10 + int(l-'0')
		}
		if l > 48 && l <= 57 {
			digit = true
			n = n*10 + int(l-'0')
		}
		if l == '-' && !digit {
			sign = -1
		}
	}
	return n * sign
}
