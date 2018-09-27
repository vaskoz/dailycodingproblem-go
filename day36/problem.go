package day36

// BST represents a binary search tree.
type BST struct {
	val         int
	left, right *BST
}

// SecondLargest takes the root of a binary search tree and returns the
// pointer to the 2nd largest value.
// Runs in O(height) and O(1) space.
func SecondLargest(root *BST) int {
	var prev, cur *BST
	if root.right != nil {
		prev = root
		cur = root.right
		for cur.right != nil {
			cur = cur.right
			prev = prev.right
		}
		cur = prev
	} else {
		cur = root.left
		for cur.right != nil {
			cur = cur.right
		}
	}
	return cur.val
}
