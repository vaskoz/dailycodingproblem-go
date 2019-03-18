package day208

// SinglyLL is a singly linked list of integers.
type SinglyLL struct {
	Value int
	Next  *SinglyLL
}

// PartitionSinglyLL partitions the singly linked list of
// integers around the pivot point 'k'.
// Runs in O(N) time and O(N) space.
func PartitionSinglyLL(head *SinglyLL, k int) *SinglyLL {
	leftHead := &SinglyLL{}
	rightHead := &SinglyLL{}
	left := leftHead
	right := rightHead
	for head != nil {
		if head.Value < k {
			left.Next = &SinglyLL{Value: head.Value}
			left = left.Next
		} else {
			right.Next = &SinglyLL{Value: head.Value}
			right = right.Next
		}
		head = head.Next
	}
	left.Next = rightHead.Next
	return leftHead.Next
}
