package day422

// BinaryTree is an integer binary tree.
type BinaryTree struct {
	Val         int
	Left, Right *BinaryTree
}

// MergeBinaryTrees returns a new binary tree that is the "merge"
// of two other binary trees.
// Merge is defined as the sum of equivalently placed nodes.
// Runs in O(N) time.
func MergeBinaryTrees(one, two *BinaryTree) *BinaryTree {
	switch {
	case one == nil && two == nil:
		return nil
	case one == nil:
		return &BinaryTree{two.Val,
			MergeBinaryTrees(nil, two.Left),
			MergeBinaryTrees(nil, two.Right),
		}
	case two == nil:
		return &BinaryTree{one.Val,
			MergeBinaryTrees(one.Left, nil),
			MergeBinaryTrees(one.Right, nil),
		}
	default:
		return &BinaryTree{one.Val + two.Val,
			MergeBinaryTrees(one.Left, two.Left),
			MergeBinaryTrees(one.Right, two.Right),
		}
	}
}
