package day296

// BST is a binary search tree of integers.
type BST struct {
	Value       int
	Left, Right *BST
}

// SortedSliceToBST constructs a BST from a sorted slice of ints.
func SortedSliceToBST(sorted []int) *BST {
	if len(sorted) == 0 {
		return nil
	}
	mid := len(sorted) / 2
	root := &BST{Value: sorted[mid]}
	root.Left = SortedSliceToBST(sorted[:mid])
	root.Right = SortedSliceToBST(sorted[mid+1:])
	return root
}
