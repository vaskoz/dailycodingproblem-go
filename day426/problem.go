package day426

// BinaryTree is a binary tree of integers.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// MinimumSumLevel returns both the 1-based level and sum of that level.
// Runs in O(N) time and O(1) extra space.
// If you pass in "nil", the sum is 0 and the level is 0.
func MinimumSumLevel(tree *BinaryTree) (minSum int, minLevel int) {
	if tree == nil {
		return
	}

	minSum = int(^uint(0) >> 1)
	minLevel = minSum
	current := make([]*BinaryTree, 0, 1)
	current = append(current, tree)
	level := 1

	for len(current) != 0 {
		nextLevel := make([]*BinaryTree, 0, 2*cap(current))

		var sum int
		for _, ptr := range current {
			sum += ptr.Value

			if ptr.Left != nil {
				nextLevel = append(nextLevel, ptr.Left)
			}

			if ptr.Right != nil {
				nextLevel = append(nextLevel, ptr.Right)
			}
		}

		if sum < minSum {
			minSum = sum
			minLevel = level
		}
		level++

		current = nextLevel
	}

	return minSum, minLevel
}
