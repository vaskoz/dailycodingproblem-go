package day247

// BinaryTree is a simple binary tree without a value.
type BinaryTree struct {
	// Value int // is unnecessary. We don't need any value.
	Left, Right *BinaryTree
}

// IsHeightBalanced returns true if the tree is height balanced.
// It returns false otherwise.
// Runs in O(N) time.
func IsHeightBalanced(root *BinaryTree) bool {
	_, balanced := isHeightBalanced(root)
	return balanced
}

func isHeightBalanced(n *BinaryTree) (int, bool) {
	if n == nil {
		return 0, true
	}
	leftHeight, leftBalance := isHeightBalanced(n.Left)
	if !leftBalance {
		return 0, false
	}
	rightHeight, rightBalance := isHeightBalanced(n.Right)
	if !rightBalance {
		return 0, false
	}
	if diff := abs(leftHeight - rightHeight); diff > 1 {
		return 0, false
	}
	return max(leftHeight, rightHeight) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
