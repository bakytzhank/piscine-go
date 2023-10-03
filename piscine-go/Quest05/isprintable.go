package piscine

func IsPrintable(s string) bool {
	for _, l := range s {
		if l < 32 || l > 126 {
			return false
		}
	}
	return true
}
