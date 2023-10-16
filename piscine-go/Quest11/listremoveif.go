package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var previous *NodeL = nil
	current := l.Head

	for current != nil {
		if current.Data == data_ref {
			if previous != nil {
				previous.Next = current.Next
				if current.Next == nil {
					l.Tail = previous
				}
			} else {
				l.Head = current.Next
				if current.Next == nil {
					l.Tail = nil
				}
			}
		} else {
			previous = current
		}
		current = current.Next
	}
}
