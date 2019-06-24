package day305

// LL is a linked list of integers.
type LL struct {
	Value int
	Next  *LL
}

// RemoveConsecutiveSumZero removes all sublists that sum to zero.
// Runs in O(N^2) time and O(1) space.
func RemoveConsecutiveSumZero(head *LL) *LL {
	fakeHead := &LL{0, head}
	for start := fakeHead; start != nil; start = start.Next {
		var sum int
		var farthest *LL
		var found bool
		for next := start.Next; next != nil; next = next.Next {
			sum += next.Value
			if sum == 0 {
				farthest = next.Next
				found = true
			}
		}
		if found {
			start.Next = farthest
		}
	}
	return fakeHead.Next
}
