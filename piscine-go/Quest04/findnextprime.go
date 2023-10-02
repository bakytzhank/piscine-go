package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	prime := nb - 1
	found := false
	for !found {
		prime++
		if IsPrime(prime) {
			found = true
		}
	}
	return prime
}
