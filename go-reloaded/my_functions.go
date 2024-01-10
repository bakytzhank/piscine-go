package cadet

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

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func Capitalize(s string) string {
	result := []byte(s)
	upperNext := true
	for i := 0; i < len(result); i++ {
		if isAlphanumeric(result[i]) {
			if upperNext {
				if result[i] >= 'a' && result[i] <= 'z' {
					result[i] -= 32 // Convert to uppercase
				}
				upperNext = false
			} else if result[i] >= 'A' && result[i] <= 'Z' {
				result[i] += 32 // Convert to lowercase
			}
		} else {
			upperNext = true
		}
	}
	return string(result)
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}

func IndexBase(r rune, baseRunes []rune) int {
	for i, l := range baseRunes {
		if l == r {
			return i
		}
	}
	return 0
}

func toDeci(s string, base string) int {
	n := len(base)
	rev := ""
	for _, l := range s {
		rev = string(l) + rev
	}
	nbr := []rune(rev)
	baseRunes := []rune(base)
	res := 0
	for i := 0; i < len(nbr); i++ {
		res = res + IndexBase(nbr[i], baseRunes)*RecursivePower(n, i)
	}
	return res
}