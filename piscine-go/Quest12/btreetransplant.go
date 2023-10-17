package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent.Data > node.Data {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	rplc.Parent = node.Parent
	return root
}
