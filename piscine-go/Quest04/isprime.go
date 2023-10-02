package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	i, j := 5, 2
	for i*i <= nb {
		if nb%i == 0 {
			return false
		}
		i += j
		j = 6 - j
	}
	return true
}
