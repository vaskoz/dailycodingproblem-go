package day465

// SinglyLL represents a node in a singly linked list.
type SinglyLL struct {
	Value interface{}
	Next  *SinglyLL
}

// ReverseSinglyLinkedList reverses a singly linked list in O(1) space
// and O(N) time.
func ReverseSinglyLinkedList(head *SinglyLL) *SinglyLL {
	var reversed *SinglyLL

	for head != nil {
		next := head.Next
		head.Next = reversed
		reversed = head
		head = next
	}

	return reversed
}
