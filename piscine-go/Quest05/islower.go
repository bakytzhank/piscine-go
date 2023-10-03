package piscine

func IsLower(s string) bool {
	for _, l := range s {
		if l < 97 || l > 122 {
			return false
		}
	}
	return true
}
