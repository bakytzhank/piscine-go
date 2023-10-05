package piscine

func MakeRange(min, max int) []int {
	var resRange []int
	if min < max {
		resRange = make([]int, max-min)
		for i, j := min, 0; i < max; i++ {
			resRange[j] = i
			j++
		}
	}
	return resRange
}
