package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	var prev *NodeI = nil
	newNode := &NodeI{Data: data_ref}
	current := l
	if data_ref < l.Data {
		newNode.Next = l
		l = newNode
	} else {
		prev = current
		current = current.Next
		for current != nil {
			if data_ref < current.Data {
				newNode.Next = current
				prev.Next = newNode
				return l
			}
			prev = current
			current = current.Next
		}
		newNode.Next = current
		prev.Next = newNode
	}
	return l
}
