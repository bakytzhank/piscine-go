package piscine

func IndexInBase(r rune, baseRunes []rune) int {
	for i, l := range baseRunes {
		if l == r {
			return i
		}
	}
	return 0
}

func AtoiBase(s string, base string) int {
	n := len(base)
	minusplus := false
	for _, char := range base {
		if char == '+' || char == '-' {
			minusplus = true
		}
	}
	if n < 2 || minusplus {
		return 0
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	rev := ""
	for _, l := range s {
		rev = string(l) + rev
	}
	nbr := []rune(rev)
	baseRunes := []rune(base)
	res := 0
	for i := 0; i < len(nbr); i++ {
		res = res + IndexInBase(nbr[i], baseRunes)*RecursivePower(n, i)
	}
	return res
}
