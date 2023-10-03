package piscine

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
