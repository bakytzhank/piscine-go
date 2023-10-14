package piscine

func Rot14(s string) string {
	resString := ""
	var offset rune
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			offset = rune(14)
			if ch >= 'a' && ch <= 'z' && ch+offset > 'z' || ch >= 'A' && ch <= 'Z' && ch+offset > 'Z' {
				offset = rune(-12)
			}
			resString += string(ch + offset)
		} else {
			resString += string(ch)
		}
	}
	return resString
}
