package day78

// SinglyLL is a singly linked list of integers.
type SinglyLL struct {
	Value int
	Next  *SinglyLL
}

// MergeKSortedLists merges k-sorted singly linked lists.
func MergeKSortedLists(lists []*SinglyLL) *SinglyLL {
	var combined, ptr *SinglyLL
	done := false
	for !done {
		remaining := 0
		var smallest *SinglyLL
		smallestIndex := 0
		for i := range lists {
			if lists[i] != nil {
				if smallest == nil {
					smallest = lists[i]
					smallestIndex = i
				} else if lists[i].Value < smallest.Value {
					smallest = lists[i]
					smallestIndex = i
				}
				remaining++
			}
		}
		if remaining == 0 {
			done = true
		} else if combined == nil {
			lists[smallestIndex] = lists[smallestIndex].Next
			combined = smallest
			ptr = smallest
			ptr.Next = nil
		} else {
			lists[smallestIndex] = lists[smallestIndex].Next
			ptr.Next = smallest
			ptr = ptr.Next
			ptr.Next = nil
		}
	}
	return combined
}
