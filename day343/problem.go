package day343

// BST is an integer binary search tree.
type BST struct {
	Val         int
	Left, Right *BST
}

// SumBSTRange takes a BST root node and
// two values indicating a range.
// Runs in linear time recursively.
func SumBSTRange(root *BST, a, b int) int {
	if root == nil || a > b {
		return 0
	}
	leftSum := SumBSTRange(root.Left, a, root.Val)
	rightSum := SumBSTRange(root.Right, root.Val, b)
	if root.Val >= a && root.Val <= b {
		return root.Val + leftSum + rightSum
	}
	return leftSum + rightSum
}
