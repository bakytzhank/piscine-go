package piscine

func inOrderTraversal(root *TreeNode, prev *string) bool {
	if root != nil {
		if !inOrderTraversal(root.Left, prev) {
			return false
		}
		if *prev != "" && root.Data <= *prev {
			return false
		}
		*prev = root.Data
		return inOrderTraversal(root.Right, prev)
	}
	return true
}

func BTreeIsBinary(root *TreeNode) bool {
	var prev string
	return inOrderTraversal(root, &prev)
}
