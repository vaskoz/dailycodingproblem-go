package day204

// CompleteBinaryTree is a complete binary tree.
type CompleteBinaryTree struct {
	Value       interface{}
	Left, Right *CompleteBinaryTree
}

// CountCompleteBinaryTree returns the number of nodes in a complete
// binary tree is less than O(N) time.
// Runs in O(height^2) time.
func CountCompleteBinaryTree(tree *CompleteBinaryTree) int {
	if tree == nil {
		return 0
	}
	var leftHeight, rightHeight int
	for left := tree; left != nil; left = left.Left {
		leftHeight++
	}
	for right := tree; right != nil; right = right.Right {
		rightHeight++
	}
	if leftHeight == rightHeight {
		return (1 << uint(leftHeight)) - 1
	}
	return CountCompleteBinaryTree(tree.Left) + CountCompleteBinaryTree(tree.Right) + 1
}
