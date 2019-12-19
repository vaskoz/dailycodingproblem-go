package day394

// IntBinaryTree is a binary tree of integers.
type IntBinaryTree struct {
	Value       int
	Left, Right *IntBinaryTree
}

// RootToLeafTarget answers true/false if there is a path
// that sums to target from root to leaf.
func RootToLeafTarget(root *IntBinaryTree, target int) bool {
	if root.Left == nil && root.Right == nil {
		return target == root.Value
	}

	if root.Left != nil {
		if ans := RootToLeafTarget(root.Left, target-root.Value); ans {
			return ans
		}
	}

	if root.Right != nil {
		if ans := RootToLeafTarget(root.Right, target-root.Value); ans {
			return ans
		}
	}

	return false
}
