package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		str += strs[i]
		if i != len(strs)-1 {
			str += sep
		}
	}
	return str
}
