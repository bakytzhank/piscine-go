package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	result := 0
	baseLen := len(base)

	for i := range s {
		digit := getDigit(base, s[i])
		result = result*baseLen + digit
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}

		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func getDigit(base string, char byte) int {
	for i := range base {
		if base[i] == char {
			return i
		}
	}
	return -1
}
