package piscine

func IsUpper(s string) bool {
	for _, l := range s {
		if l < 65 || l > 90 {
			return false
		}
	}
	return true
}
