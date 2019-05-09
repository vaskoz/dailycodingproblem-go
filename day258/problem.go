package day258

// IntBinaryTree is a binary tree of integers.
type IntBinaryTree struct {
	Value       int
	Left, Right *IntBinaryTree
}

// BoustrophedonOrder works as follows.
// https://en.wikipedia.org/wiki/Boustrophedon
func BoustrophedonOrder(root *IntBinaryTree) []int {
	if root == nil {
		return nil
	}
	var result []int
	levels := []*IntBinaryTree{root}
	leftToRight := true
	for len(levels) != 0 {
		var nextLevel []*IntBinaryTree
		for _, n := range levels {
			result = append(result, n.Value)
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}
		if !leftToRight {
			toReverse := result[len(result)-len(levels):]
			for i := 0; i < len(toReverse)/2; i++ {
				toReverse[i], toReverse[len(toReverse)-1-i] = toReverse[len(toReverse)-1-i], toReverse[i]
			}
		}
		leftToRight = !leftToRight
		levels = nextLevel
	}
	return result
}
