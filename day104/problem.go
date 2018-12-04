package day104

// DoublyLL is a doubly linked list.
type DoublyLL struct {
	Value      int
	Prev, Next *DoublyLL
}

// SinglyLL is a singly linked list.
type SinglyLL struct {
	Value int
	Next  *SinglyLL
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

// IsPalindromeSinglyLLStack answers the question "is this list a palindrome?".
// Runs on a singly linked list.
// Runs in O(N) time and O(N) space.
func IsPalindromeSinglyLLStack(head *SinglyLL) bool {
	var n int
	var data []int
	for tail := head; tail != nil; tail = tail.Next {
		data = append(data, tail.Value)
		n++
	}
	for i := range data {
		if head == nil || head.Value != data[len(data)-1-i] {
			return false
		}
		head = head.Next
	}
	return head == nil
}

// IsPalindromeSinglyLLReverse answers the question "is this list a palindrome?".
// Runs on a singly linked list.
// Runs in O(N) time and O(1) space.
// NOTE: This modifies the linked list during execution, but restores it after completion.
func IsPalindromeSinglyLLReverse(head *SinglyLL) bool {
	var n int
	for tail := head; tail != nil; tail = tail.Next {
		n++
	}
	halfPoint := head
	for i := 0; i < (n/2)-1; i++ {
		halfPoint = halfPoint.Next
	}
	if n%2 == 1 {
		halfPoint = halfPoint.Next
	}
	reversed := ReverseSinglyLL(halfPoint.Next)
	halfPoint.Next = reversed
	isPalindrome := true
	half := halfPoint.Next
	for i := 0; i < n/2; i++ {
		if half.Value != head.Value {
			isPalindrome = false
			break
		}
		half = half.Next
		head = head.Next
	}
	halfPoint.Next = ReverseSinglyLL(halfPoint.Next)
	return isPalindrome
}

// ReverseSinglyLL reverses a singly linked list in O(1) space
// and O(N) time.
func ReverseSinglyLL(head *SinglyLL) *SinglyLL {
	var reversed *SinglyLL
	for head != nil {
		next := head.Next
		head.Next = reversed
		reversed = head
		head = next
	}
	return reversed
}
