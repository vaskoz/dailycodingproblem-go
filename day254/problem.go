package day254

// BinaryTree with an int value and a left and right pointer.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// ConvertBinaryTreeToFull deletes nodes in a binary tree
// until it's a full tree.
func ConvertBinaryTreeToFull(root *BinaryTree) *BinaryTree {
	if root.Left == nil && root.Right == nil {
		return root
	}
	var left, right *BinaryTree
	if root.Left != nil {
		left = ConvertBinaryTreeToFull(root.Left)
	}
	if root.Right != nil {
		right = ConvertBinaryTreeToFull(root.Right)
	}
	if left != nil && right != nil {
		root.Left = left
		root.Right = right
		return root
	} else if left != nil {
		return left
	}
	return right
}
