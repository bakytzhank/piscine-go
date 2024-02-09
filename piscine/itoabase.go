package piscine

func ItoaBase(value, base int) string {
	if value == -9223372036854775808 {
		return "-" + ItoaBase(-(value/base), base) + string("0123456789ABCDEF"[-(value%base)])
	}

	if value == 0 {
		return "0"
	}

	baseStr := "0123456789ABCDEF"
	result := ""
	isNegative := false

	if value < 0 {
		isNegative = true
		value = -value
	}

	for value > 0 {
		remainder := value % base
		result = string(baseStr[remainder]) + result
		value /= base
	}

	if isNegative {
		result = "-" + result
	}

	return result
}