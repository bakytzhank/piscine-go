package piscine

func BTreeLevelCount(root *TreeNode) int {
	var level int
	if root == nil {
		return 0
	}
	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)
	if leftHeight > rightHeight {
		level = leftHeight
	} else {
		level = rightHeight
	}
	return level + 1
}
