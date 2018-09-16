package day26

// Node is a node in a singly linked list.
type Node struct {
	Val  int
	Next *Node
}

// RemoveFromEnd deletes the kth element from the end of the list.
// Loops the list once O(N) and uses constant space O(1)
func RemoveFromEnd(head *Node, kth int) *Node {
	originalHead := head
	kptr := head
	passedHead := false
	for i := 0; i <= kth; i++ {
		if head == nil {
			passedHead = true
		} else {
			head = head.Next
		}
	}
	for head != nil {
		head = head.Next
		kptr = kptr.Next
	}
	if passedHead {
		kptr = kptr.Next
		return kptr
	}
	kptr.Next = kptr.Next.Next
	return originalHead
}
