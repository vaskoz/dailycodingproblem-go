package day20

// Node is a singly linked list. NOTE: Doubly-linked list won't work here with shared paths.
type Node struct {
	Value interface{}
	Next  *Node
}

// FindCommonNode finds the common element in a linked list in O(N+M) time and O(1) space.
func FindCommonNode(one, two *Node) *Node {
	var oneLen, twoLen int
	for e := one; e != nil; e = e.Next {
		oneLen++
	}
	for e := two; e != nil; e = e.Next {
		twoLen++
	}
	var longer, shorter *Node
	var diff int
	if oneLen > twoLen {
		longer = one
		shorter = two
		diff = oneLen - twoLen
	} else {
		longer = two
		shorter = one
		diff = twoLen - oneLen
	}
	for i := 0; i < diff; i++ {
		longer = longer.Next
	}
	for longer != nil {
		if longer == shorter {
			return longer
		}
		longer = longer.Next
		shorter = shorter.Next
	}
	return nil
}
