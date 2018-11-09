package deepest

// BinaryTree represents a tree with a left and right subtree.
type BinaryTree struct {
	// val         int // value doesn't matter here.
	left, right *BinaryTree
}

// CountNodes tallies the total number of nodes in the binary tree.
func CountNodes(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	return CountNodes(root.left) + CountNodes(root.right) + 1
}
