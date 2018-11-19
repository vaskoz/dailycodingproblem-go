package day89

// BST represents a binary search tree.
type BST struct {
	Value       int
	Left, Right *BST
}

// IsValidBST checks if the given tree conforms to the BST property.
func IsValidBST(root *BST) bool {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	return isValidBST(root.Left, MinInt, root.Value) && isValidBST(root.Right, root.Value, MaxInt)
}

func isValidBST(root *BST, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Value > min && root.Value <= max {
		return isValidBST(root.Left, min, root.Value) && isValidBST(root.Right, root.Value, max)
	}
	return false
}
