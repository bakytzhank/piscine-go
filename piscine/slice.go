package piscine

func Slice(a []string, nbrs... int) []string{
	var start, end int
	switch len(nbrs) {
	case 1:
		start = nbrs[0]
		end = len(a)
	case 2:
		start = nbrs[0]
		end = nbrs[1]
	}

	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		end = len(a) + end
	}

	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}

	if start > end {
		return nil
	}

	if end > len(a) {
		end = len(a)
	}

	return a[start:end]
}
