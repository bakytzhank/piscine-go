package piscine

func ListClear(l *List) {
	// current := l.Head
	l.Head = nil
	l.Tail = nil
	// for current != nil {
	// 	next := current.Next
	// 	current.Next = nil
	// 	current = next
	// }
}
