package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		}

		temp := BTreeMin(root.Right)
		root.Data = temp.Data
		root.Right = BTreeDeleteNode(root.Right, temp)
	}
	return root
}
