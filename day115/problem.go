package day115

// IntBinaryTree is a binary tree of integers.
type IntBinaryTree struct {
	Value       int
	Left, Right *IntBinaryTree
}

// SameTreeSubstructure returns true if 't' exists somewhere inside of 's'.
func SameTreeSubstructure(t, s *IntBinaryTree) bool {
	if s == nil {
		return false
	}
	return equalTrees(t, s) || SameTreeSubstructure(t, s.Left) || SameTreeSubstructure(t, s.Right)
}

func equalTrees(one, two *IntBinaryTree) bool {
	if one == nil && two == nil {
		return true
	} else if (one == nil && two != nil) || (one != nil && two == nil) || (one.Value != two.Value) {
		return false
	}
	return equalTrees(one.Left, two.Left) || equalTrees(one.Right, two.Right)
}
