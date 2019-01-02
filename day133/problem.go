package day133

// BST is a binary search tree of integers.
type BST struct {
	Value               int
	Left, Right, Parent *BST
}

// InorderSuccessor returns a pointer to the next bigger element.
func InorderSuccessor(node *BST) *BST {
	if node == nil {
		return nil
	} else if node.Right == nil {
		return node.Parent
	}
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node
}
