package day110

// IntBinaryTree is a binary tree with values that are integers.
type IntBinaryTree struct {
	Value       int
	Left, Right *IntBinaryTree
}

// AllPathsToLeaves return the nodes encountered from the root
// to each leave.
func AllPathsToLeaves(tree *IntBinaryTree) [][]int {
	return pathHolder(tree, []int{})
}

func pathHolder(tree *IntBinaryTree, path []int) [][]int {
	path = append(path, tree.Value)
	if tree.Left == nil && tree.Right == nil {
		leavePath := make([]int, len(path))
		copy(leavePath, path)
		return [][]int{leavePath}
	}
	var result [][]int
	if tree.Left != nil {
		leftSubtree := pathHolder(tree.Left, path)
		result = append(result, leftSubtree...)
	}
	if tree.Right != nil {
		rightSubtree := pathHolder(tree.Right, path)
		result = append(result, rightSubtree...)
	}
	return result
}
