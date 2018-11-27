package day94

// BinaryTree is a binary tree with an integer value.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// MaxPathSum returns the maximum path sum in a given
// binary tree.
func MaxPathSum(tree *BinaryTree) int {
	const MinInt = -int(^uint(0)>>1) - 1
	result := MinInt
	maxPathSum(tree, &result)
	return result
}

func maxPathSum(node *BinaryTree, res *int) int {
	if node == nil {
		return 0
	} else if node.Left == nil && node.Right == nil {
		return node.Value
	}
	leftSum := maxPathSum(node.Left, res)
	rightSum := maxPathSum(node.Right, res)
	if node.Left != nil && node.Right != nil {
		*res = max(*res, leftSum+rightSum+node.Value)
		return max(leftSum, rightSum) + node.Value
	}
	if node.Left == nil {
		return rightSum + node.Value
	}
	return leftSum + node.Value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
