package day169

// SinglyLL is a singly linked list of integers.
type SinglyLL struct {
	Value int
	Next  *SinglyLL
}

// MergesortSinglyLinkedList runs in O(N log N) time and O(1) space.
func MergesortSinglyLinkedList(head *SinglyLL) *SinglyLL {
	var size int
	ptr := head
	for ptr != nil {
		size++
		ptr = ptr.Next
	}
	if size == 1 {
		return head
	}
	ptr = head
	for i := 0; i < (size/2)-1; i++ {
		ptr = ptr.Next
	}
	half := ptr.Next
	ptr.Next = nil
	a := MergesortSinglyLinkedList(head)
	b := MergesortSinglyLinkedList(half)
	return mergeSorted(a, b)
}

func mergeSorted(a, b *SinglyLL) *SinglyLL {
	result := &SinglyLL{}
	ptr := result
	for a != nil || b != nil {
		switch {
		case a != nil && b != nil:
			if a.Value < b.Value {
				ptr.Next = a
				a = a.Next
				ptr = ptr.Next
				ptr.Next = nil
			} else {
				ptr.Next = b
				b = b.Next
				ptr = ptr.Next
				ptr.Next = nil
			}
		case a != nil:
			ptr.Next = a
			a = a.Next
			ptr = ptr.Next
			ptr.Next = nil
		case b != nil:
			ptr.Next = b
			b = b.Next
			ptr = ptr.Next
			ptr.Next = nil
		}
	}
	return result.Next
}
