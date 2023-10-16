package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l

	for i := 0; i < pos && current != nil; i++ {
		current = current.Next
	}

	return current
}
