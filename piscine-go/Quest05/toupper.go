package piscine

func ToUpper(s string) string {
	str := ""
	for _, l := range s {
		if l >= 97 && l <= 122 {
			l = rune(l) - 32
		}
		str = str + string(l)
	}
	return str
}
