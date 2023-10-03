package piscine

func IsNumeric(s string) bool {
	for _, l := range s {
		if l < 48 || l > 57 {
			return false
		}
	}
	return true
}
