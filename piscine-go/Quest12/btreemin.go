package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
