package day135

// BinaryTree is a binary tree.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// MinPathSum returns the sum of the minimum sum path from root to leaf.
func MinPathSum(head *BinaryTree) int {
	switch {
	case head.Left == nil && head.Right == nil:
		return head.Value
	case head.Left != nil && head.Right != nil:
		left := head.Value + MinPathSum(head.Left)
		right := head.Value + MinPathSum(head.Right)
		return min(left, right)
	case head.Left != nil:
		return head.Value + MinPathSum(head.Left)
	default:
		return head.Value + MinPathSum(head.Right)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
