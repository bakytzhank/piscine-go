package piscine

func ListSize(l *List) int {
	length := 0
	current := l.Head

	for current != nil {
		length++
		current = current.Next
	}

	return length
}
