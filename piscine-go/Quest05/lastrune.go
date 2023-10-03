package piscine

func LastRune(s string) rune {
	lastrune := []rune(s)
	return lastrune[len(lastrune)-1]
}
