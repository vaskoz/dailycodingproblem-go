package day145

// LL is a singly linked list of integers.
type LL struct {
	Value int
	Next  *LL
}

// SwapEveryTwo swaps every two nodes in a linked list.
func SwapEveryTwo(head *LL) *LL {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	result.Next, head.Next = head, head.Next.Next
	result.Next.Next = SwapEveryTwo(result.Next.Next)
	return result
}
