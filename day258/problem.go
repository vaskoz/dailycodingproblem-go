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
		if leftToRight {
			for _, n := range levels {
				result = append(result, n.Value)
				if n.Left != nil {
					nextLevel = append(nextLevel, n.Left)
				}
				if n.Right != nil {
					nextLevel = append(nextLevel, n.Right)
				}
			}
		} else {
			for i := len(levels) - 1; i >= 0; i-- {
				result = append(result, levels[i].Value)
				j := len(levels) - 1 - i
				if levels[j].Left != nil {
					nextLevel = append(nextLevel, levels[j].Left)
				}
				if levels[j].Right != nil {
					nextLevel = append(nextLevel, levels[j].Right)
				}
			}
		}
		levels = nextLevel
		leftToRight = !leftToRight
	}
	return result
}
