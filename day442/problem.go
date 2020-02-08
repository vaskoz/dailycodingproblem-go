package day442

// IntBinaryTree is a typical binary tree
// that stores integers for a value.
type IntBinaryTree struct {
	Value       int
	Left, Right *IntBinaryTree
}

// CartesianTree returns the corresponding Cartesian tree
// based on the given sequence.
func CartesianTree(s []int) *IntBinaryTree {
	if len(s) == 0 {
		return nil
	}

	minIndex := 0

	for i := range s {
		if s[i] < s[minIndex] {
			minIndex = i
		}
	}

	root := &IntBinaryTree{s[minIndex], nil, nil}
	root.Left = CartesianTree(s[:minIndex])
	root.Right = CartesianTree(s[minIndex+1:])

	return root
}
