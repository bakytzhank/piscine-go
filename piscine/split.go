package piscine

func FoundSep(s, sep string, index int) bool {
	if s[index:index+len(sep)] == sep {
		return true
	}
	return false
}

func Split(s, sep string) []string {
	var resSplit []string
	word := ""
	for i:=0; i<len(s); i++ {
		if i < len(s) - len(sep) {
			if FoundSep(s, sep, i) {
				resSplit = append(resSplit, word)
				word = ""
				i += (len(sep) - 1)
			} else {
				word += string(s[i])
			}
		} else {
			word += string(s[i])
		}
	}
	resSplit = append(resSplit, word)
	return resSplit
}
