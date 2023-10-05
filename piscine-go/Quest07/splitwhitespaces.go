package piscine

func SplitWhiteSpaces(s string) []string {
	var resSplit []string
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			if s[i-1] == ' ' || s[i-1] == '\n' || s[i-1] == '\t' {
				continue
			} else {
				resSplit = append(resSplit, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	resSplit = append(resSplit, word)
	return resSplit
}
