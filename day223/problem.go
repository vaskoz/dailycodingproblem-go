package day223

// BinaryTree is a simple binary tree.
type BinaryTree struct {
	Value       interface{}
	Left, Right *BinaryTree
}

// InorderMorrisTraversal performs a Morris Inorder traversal.
// Uses O(1) extra space.
func InorderMorrisTraversal(root *BinaryTree) []interface{} {
	var inorder []interface{}
	curr := root
	for curr != nil {
		if curr.Left == nil {
			inorder = append(inorder, curr.Value)
			curr = curr.Right
		} else {
			pre := curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = curr
				curr = curr.Left
			} else {
				pre.Right = nil
				inorder = append(inorder, curr.Value)
				curr = curr.Right
			}
		}
	}
	return inorder
}
