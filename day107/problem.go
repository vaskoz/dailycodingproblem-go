package day107

// BinaryTree is a binary tree of integers.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// BinaryTreeByLevel prints the binary tree level-wise.
// Runs in O(N^2) time.
func BinaryTreeByLevel(tree *BinaryTree) []int {
	var result []int
	for i := 0; i < Height(tree); i++ {
		result = append(result, GetLevel(tree, i)...)
	}
	return result
}

// GetLevel returns all the nodes at the same height from left to right.
func GetLevel(tree *BinaryTree, level int) []int {
	if tree == nil {
		return nil
	}
	if level == 0 {
		return []int{tree.Value}
	}
	var result []int
	result = append(result, GetLevel(tree.Left, level-1)...)
	result = append(result, GetLevel(tree.Right, level-1)...)
	return result
}

// Height returns the height of a binary tree.
func Height(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}
	return 1 + max(Height(tree.Left), Height(tree.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
