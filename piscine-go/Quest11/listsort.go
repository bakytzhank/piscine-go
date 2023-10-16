package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
	if current == nil {
		return nil
	} else {
		for current != nil {
			next := current.Next
			for next != nil {
				if current.Data > next.Data {
					temp := current.Data
					current.Data = next.Data
					next.Data = temp
				}
				next = next.Next
			}
			current = current.Next
		}
	}
	return l
}
