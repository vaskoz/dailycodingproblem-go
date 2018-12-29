package day125

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
	leftResult := inorder(head, current.Left, k, result)
	if len(leftResult) == 2 {
		return leftResult
	}
	node := searchBST(head, k-current.Val)
	if node != nil && node != current {
		return []int{current.Val, node.Val}
	}
	rightResult := inorder(head, current.Right, k, result)
	if len(rightResult) == 2 {
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
		// } else if target < head.Val { // this part of searchBST is unused due to the
		// 	return searchBST(head.Left, target) // inorder traversal that calls this search.
	}
	return searchBST(head.Right, target)
}
