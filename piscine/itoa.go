package piscine

func Itoa(n int) string {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	result := ""
	for n > 0 {
		digit := n % 10
		result = string(digit+'0') + result
		n /= 10
	}

	if isNegative {
		result = "-" + result
	}

	if result == "" {
		result = "0"
	}

	return result
}
