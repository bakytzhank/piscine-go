package piscine

func Compact(ptr *[]string) int {
	var resStr []string
	for _, str := range *ptr {
		if str != "" {
			resStr = append(resStr, str)
		}
	}
	*ptr = resStr
	return len(resStr)
}
