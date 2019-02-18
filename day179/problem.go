package day179

// BinaryTree is a binary tree of integers.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// BuildTreeFromPostorder reconstructs a binary tree from
// a postorder traversal.
func BuildTreeFromPostorder(lst []int) *BinaryTree {
	if len(lst) == 0 {
		return nil
	}
	root := &BinaryTree{lst[len(lst)-1], nil, nil}
	lst = lst[:len(lst)-1]
	i := len(lst) - 1
	for ; i >= 0 && lst[i] > root.Value; i-- {
	}
	root.Left = BuildTreeFromPostorder(lst[:i+1])
	root.Right = BuildTreeFromPostorder(lst[i+1:])
	return root
}
