package piscine

func CollatzCountdown(start int) int {
	if start < 1 {
		return -1
	}
	count := 0
	for start != 1 {
		if start%2 == 0 {
			start /= 2
			count++
		} else {
			start = 3*start + 1
			count++
		}
	}
	return count
}
