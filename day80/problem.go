package day80

// BinaryTree represents a tree with a left and right subtree.
type BinaryTree struct {
	// val         int // value doesn't matter here.
	left, right *BinaryTree
}

// Deepest returns the deepest node in the binary tree.
func Deepest(root *BinaryTree) (*BinaryTree, int) {
	if root != nil && root.left == nil && root.right == nil {
		return root, 1
	}
	if root.left == nil {
		n, count := Deepest(root.right)
		return n, count + 1
	} else if root.right == nil {
		n, count := Deepest(root.left)
		return n, count + 1
	}
	ln, lcount := Deepest(root.left)
	rn, rcount := Deepest(root.right)
	if lcount > rcount {
		return ln, lcount + 1
	}
	return rn, rcount + 1
}
