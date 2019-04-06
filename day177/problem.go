package day177

// LL is a singly linked list.
type LL struct {
	Value interface{}
	Next  *LL
}

// RotateRightLL rotates the linked list to the right
// by k elements.
// Runs in O(N) time and O(1) extra space.
func RotateRightLL(head *LL, k int) *LL {
	if head == nil {
		return nil
	}
	size := 1
	var end *LL
	for end = head; end.Next != nil; end = end.Next {
		size++
	}
	if k%size == 0 {
		return head
	}
	k %= size
	curr := head
	for i := 0; i < k-1; i++ {
		curr = curr.Next
	}
	newHead := curr.Next
	curr.Next = nil
	end.Next = head
	return newHead
}
