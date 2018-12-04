package day104

// DoublyLL is a doubly linked list.
type DoublyLL struct {
	Value      int
	Prev, Next *DoublyLL
}

// IsPalindromeDoublyLL answers the question "is this list a palindrome?".
// Runs on a doubly linked list.
// Runs in O(N) time and O(1) extra space.
func IsPalindromeDoublyLL(head *DoublyLL) bool {
	var n int
	var tail *DoublyLL
	for tail = head; tail.Next != nil; tail = tail.Next {
		n++
	}
	n++
	for i := 0; i < n/2; i++ {
		if head.Value != tail.Value {
			return false
		}
		head = head.Next
		tail = tail.Prev
	}
	return true
}
