package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	tail := n1
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = n2

	return ListSort(n1)
}
