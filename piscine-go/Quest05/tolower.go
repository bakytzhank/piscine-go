package piscine

func ToLower(s string) string {
	str := ""
	for _, l := range s {
		if l >= 65 && l <= 90 {
			l = rune(l) + 32
		}
		str = str + string(l)
	}
	return str
}
