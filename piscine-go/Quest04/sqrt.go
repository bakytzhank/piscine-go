package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	p := 0
	for i := 0; nb >= i*i; i++ {
		p = i
	}
	if p*p == nb {
		return p
	}
	return 0
}
