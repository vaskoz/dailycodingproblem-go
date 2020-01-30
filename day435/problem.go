package day435

// BinaryTree is a binary tree with rune values.
type BinaryTree struct {
	value       string
	left, right *BinaryTree
}

// Reconstruct takes a pre-order and in-order string representation
// and then returns a BinaryTree that conforms to both of those representations.
func Reconstruct(preorder, inorder []string) (*BinaryTree, []string) {
	node := &BinaryTree{preorder[0], nil, nil}
	preorder = preorder[1:]

	if len(inorder) == 1 {
		return node, preorder
	}

	index := search(inorder, node.value)
	node.left, preorder = Reconstruct(preorder, inorder[:index])
	node.right, preorder = Reconstruct(preorder, inorder[index+1:])

	return node, preorder
}

func search(strs []string, target string) int {
	var result int

	for index, str := range strs {
		if str == target {
			result = index
			break
		}
	}

	return result
}
