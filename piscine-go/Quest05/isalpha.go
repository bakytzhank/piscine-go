package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, l := range s {
		if l < 48 || (l > 57 && l < 65) || (l > 90 && l < 97) || l > 122 {
			return false
		}
	}
	return true
}
