package piscine

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

func fromDeci(nbr int, base string) string {
	n := len(base)
	Converted := ""
	if nbr == 0 {
		Converted = string(base[0])
		return Converted
	}
	var res []rune
	for nbr > 0 {
		remainder := nbr % n
		res = append(res, rune(base[remainder]))
		nbr /= n
	}
	for k := len(res) - 1; k >= 0; k-- {
		Converted += string(res[k])
	}
	return Converted
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return fromDeci(toDeci(nbr, baseFrom), baseTo)
}
