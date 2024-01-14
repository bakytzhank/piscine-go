package piscine

func IsPrime(n int) bool {
	if n < 2 {
		return false
	} 
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for i:= nb; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
