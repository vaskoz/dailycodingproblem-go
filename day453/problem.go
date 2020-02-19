package day453

// BST is an integer binary search tree.
type BST struct {
	Val         int
	Left, Right *BST
}

// FindTwoSumKInBST returns a slice of integer nodes that sum to K.
// If no two nodes exist that sum to K, then return nil.
// Runs in O(N log N) where N is the size of the BST.
func FindTwoSumKInBST(head *BST, k int) []int {
	return inorder(head, head, k, []int{})
}

func inorder(head, current *BST, k int, result []int) []int {
	if current == nil {
		return result
	}

	if leftResult := inorder(head, current.Left, k, result); len(leftResult) == 2 {
		return leftResult
	}

	if node := searchBST(head, k-current.Val); node != nil && node != current {
		return []int{current.Val, node.Val}
	}

	if rightResult := inorder(head, current.Right, k, result); len(rightResult) == 2 {
		return rightResult
	}

	return result
}

func searchBST(head *BST, target int) *BST {
	if head == nil {
		return nil
	}

	if target == head.Val {
		return head
	}

	return searchBST(head.Right, target)
}
