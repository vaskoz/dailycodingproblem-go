package day398

// SinglyLL is a singly linked list of interface{} type.
type SinglyLL struct {
	Value interface{}
	Next  *SinglyLL
}

// RemoveKthFromEnd removes the k-th node from the
// end of the list and return the head of the list.
// Runs in O(N) time and O(1) space.
func RemoveKthFromEnd(head *SinglyLL, kth int) *SinglyLL {
	headPtr := &SinglyLL{Value: nil, Next: head}
	front := headPtr
	follower := headPtr

	for i := 0; i <= kth; i++ {
		if front == nil {
			panic("less than k elements in list")
		}

		front = front.Next
	}

	for ; front != nil; front = front.Next {
		follower = follower.Next
	}

	follower.Next = follower.Next.Next

	return headPtr.Next
}
