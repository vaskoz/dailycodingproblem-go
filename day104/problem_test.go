package day104

import "testing"

var testcases = []struct {
	data         []int
	isPalindrome bool
}{
	{[]int{1, 4, 3, 4, 1}, true},
	{[]int{1, 4}, false},
}

func TestIsPalindromeDoublyLL(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		head := createDoublyLL(tc.data)
		if result := IsPalindromeDoublyLL(head); result != tc.isPalindrome {
			t.Errorf("For input %v expected %v got %v", tc.data, tc.isPalindrome, result)
		}
	}
}

func BenchmarkIsPalindromeDoublyLL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			head := createDoublyLL(tc.data)
			IsPalindromeDoublyLL(head)
		}
	}
}

func TestIsPalindromeSinglyLLStack(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		head := createSinglyLL(tc.data)
		if result := IsPalindromeSinglyLLStack(head); result != tc.isPalindrome {
			t.Errorf("For input %v expected %v got %v", tc.data, tc.isPalindrome, result)
		}
	}
}

func BenchmarkIsPalindromeSinglyLLStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			head := createSinglyLL(tc.data)
			IsPalindromeSinglyLLStack(head)
		}
	}
}

func TestIsPalindromeSinglyLLReverse(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		head := createSinglyLL(tc.data)
		if result := IsPalindromeSinglyLLReverse(head); result != tc.isPalindrome {
			t.Errorf("For input %v expected %v got %v", tc.data, tc.isPalindrome, result)
		}
	}
}

func BenchmarkIsPalindromeSinglyLLReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			head := createSinglyLL(tc.data)
			IsPalindromeSinglyLLReverse(head)
		}
	}
}

func createDoublyLL(input []int) *DoublyLL {
	if len(input) == 0 {
		return nil
	}
	head := &DoublyLL{input[0], nil, nil}
	current := head
	for i := 1; i < len(input); i++ {
		next := &DoublyLL{Value: input[i], Prev: current, Next: nil}
		current.Next = next
		current = next
	}
	return head
}

func createSinglyLL(input []int) *SinglyLL {
	if len(input) == 0 {
		return nil
	}
	head := &SinglyLL{input[0], nil}
	current := head
	for i := 1; i < len(input); i++ {
		next := &SinglyLL{input[i], nil}
		current.Next = next
		current = next
	}
	return head
}
