package piscine

func NRune(s string, n int) rune {
	nrune := []rune(s)
	if n < 1 || n > len(nrune) {
		return 0
	}
	return nrune[n-1]
}
