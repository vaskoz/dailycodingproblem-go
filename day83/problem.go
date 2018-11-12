package day83

// BinaryTree is a binary tree with string Values.
type BinaryTree struct {
	Value       string
	Left, Right *BinaryTree
}

// InvertBinaryTree works in-place on the given tree.
func InvertBinaryTree(tree *BinaryTree) {
	if tree == nil {
		return
	}
	InvertBinaryTree(tree.Left)
	InvertBinaryTree(tree.Right)
	tree.Left, tree.Right = tree.Right, tree.Left
}
