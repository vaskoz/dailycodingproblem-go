package day146

// BinaryTree contains values of either zero (0) or one (1).
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// PruneZeroSubtrees prunes all the subtrees containing only zeros.
// NOTE: This mutates the input.
func PruneZeroSubtrees(head *BinaryTree) *BinaryTree {
	if head == nil {
		return nil
	}
	head.Left = PruneZeroSubtrees(head.Left)
	head.Right = PruneZeroSubtrees(head.Right)
	if head.Value == 0 && head.Left == nil && head.Right == nil {
		return nil
	}
	return head
}
