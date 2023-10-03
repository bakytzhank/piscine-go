package piscine

func AlphaCount(s string) int {
	n := 0
	for _, l := range s {
		if l >= 65 && l <= 90 || l >= 97 && l <= 122 {
			n++
		}
	}
	return n
}
