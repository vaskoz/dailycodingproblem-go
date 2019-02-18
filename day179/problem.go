package day179

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

// BinaryTree is a binary tree of integers.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// BuildTreeFromPostorder reconstructs a binary tree from
// a postorder traversal.
// Runs in O(N^2) time.
func BuildTreeFromPostorder(lst []int) *BinaryTree {
	if len(lst) == 0 {
		return nil
	}
	root := &BinaryTree{lst[len(lst)-1], nil, nil}
	lst = lst[:len(lst)-1]
	i := len(lst) - 1
	for ; i >= 0 && lst[i] > root.Value; i-- {
	}
	root.Left = BuildTreeFromPostorder(lst[:i+1])
	root.Right = BuildTreeFromPostorder(lst[i+1:])
	return root
}

// BuildTreeFromPostorderLinear reconstructs a binary tree from
// a postorder traversal.
// Runs in O(N) time.
func BuildTreeFromPostorderLinear(lst []int) *BinaryTree {
	return buildTreeFromPostorderLinear(&lst, minInt, maxInt)
}

func buildTreeFromPostorderLinear(lst *[]int, min, max int) *BinaryTree {
	if len(*lst) == 0 {
		return nil
	}
	var root *BinaryTree
	key := (*lst)[len(*lst)-1]
	if key > min && key < max {
		root = &BinaryTree{key, nil, nil}
		*lst = (*lst)[:len(*lst)-1]
		if len(*lst) > 0 {
			root.Right = buildTreeFromPostorderLinear(lst, key, max)
			root.Left = buildTreeFromPostorderLinear(lst, min, key)
		}
	}
	return root
}
