package piscine

func ListLast(l *List) interface{} {
	current := l.Tail
	if current != nil {
		return current.Data
	}
	return nil
}
