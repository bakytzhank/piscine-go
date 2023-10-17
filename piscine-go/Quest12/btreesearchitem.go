package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || elem == root.Data {
		return root
	}
	// if res := BTreeSearchItem(root.Left, elem); res != nil {
	// 		return res
	// }
	// if res := BTreeSearchItem(root.Right, elem); res != nil {
	// 	return res
	// }
	// return nil
	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}
