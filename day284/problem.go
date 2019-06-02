package day284

// BinaryTree is a binary tree of integers.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// FindCousins returns all the cousins of a given node.
func FindCousins(root *BinaryTree, node int) []int {
	nodeLevel := level(root, node, 1)
	return levelCousins(root, node, nodeLevel, []int{})
}

func levelCousins(root *BinaryTree, node, lev int, cousins []int) []int {
	if root == nil || lev < 2 {
		return cousins
	}
	if lev == 2 {
		if (root.Left != nil && root.Left.Value == node) ||
			(root.Right != nil && root.Right.Value == node) {
			return cousins
		}
		if root.Left != nil {
			cousins = append(cousins, root.Left.Value)
		}
		if root.Right != nil {
			cousins = append(cousins, root.Right.Value)
		}
	} else if lev > 2 {
		cousins = levelCousins(root.Left, node, lev-1, cousins)
		cousins = levelCousins(root.Right, node, lev-1, cousins)
	}
	return cousins
}

func level(root *BinaryTree, node, lev int) int {
	if root == nil {
		return 0
	} else if root.Value == node {
		return lev
	} else if down := level(root.Left, node, lev+1); down != 0 {
		return down
	}
	return level(root.Right, node, lev+1)
}
