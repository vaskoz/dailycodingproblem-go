package day8

// BinaryTree represents a binary unival tree.
type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

// CountUnivalSubtrees counts the number of unival subtrees.
func CountUnivalSubtrees(root *BinaryTree) int {
	count, _ := countUnivalSubtrees(root)
	return count
}

func countUnivalSubtrees(root *BinaryTree) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, leftUnival := countUnivalSubtrees(root.left)
	right, rightUnival := countUnivalSubtrees(root.right)
	total := left + right
	if leftUnival && rightUnival {
		if root.left != nil && root.value != root.left.value {
			return total, false
		}
		if root.right != nil && root.value != root.right.value {
			return total, false
		}
		return total + 1, true
	}
	return total, false
}
