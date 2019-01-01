package day131

// RandomSinglyLL is a singly linked list with a random pointer pointing
// anywhere in the linked list.
type RandomSinglyLL struct {
	Value        int
	Next, Random *RandomSinglyLL
}

// DeepCloneRandomSinglyLL deeply clones a RandomSinglyLL.
// Works in O(N) time and O(1) extra space aside from the O(N) extra space of the result.
func DeepCloneRandomSinglyLL(head *RandomSinglyLL) *RandomSinglyLL {
	if head == nil {
		return nil
	}
	original := head
	for original != nil {
		clone := &RandomSinglyLL{Value: original.Value, Random: nil, Next: original.Next}
		original.Next = clone
		original = original.Next.Next
	}
	original = head
	for original != nil {
		if original.Random != nil {
			original.Next.Random = original.Random.Next
		}
		original = original.Next.Next
	}
	result := head.Next
	origPtr, copyPtr := head, head.Next
	for origPtr != nil {
		origPtr.Next = origPtr.Next.Next
		if copyPtr.Next == nil {
			break
		}
		origPtr = origPtr.Next
		copyPtr.Next = copyPtr.Next.Next
		copyPtr = copyPtr.Next
	}
	return result
}
