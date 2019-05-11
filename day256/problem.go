package day256

// SinglyLL is a singly linked list of integers.
type SinglyLL struct {
	Value int
	Next  *SinglyLL
}

// RearrangeAlternating alternates from low->high->low.
func RearrangeAlternating(head *SinglyLL) *SinglyLL {
	if head == nil {
		return nil
	}
	prev := head
	curr := head.Next
	for curr != nil {
		if prev.Value > curr.Value {
			prev.Value, curr.Value = curr.Value, prev.Value
		}
		if curr.Next != nil && curr.Next.Value > curr.Value {
			curr.Value, curr.Next.Value = curr.Next.Value, curr.Value
		}
		prev = curr.Next
		if curr.Next == nil {
			break
		}
		curr = curr.Next.Next
	}
	return head
}
