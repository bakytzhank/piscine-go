package piscine

func AppendRange(min, max int) []int {
	var resRange []int
	if min < max {
		for i := min; i < max; i++ {
			resRange = append(resRange, i)
		}
	}
	return resRange
}
