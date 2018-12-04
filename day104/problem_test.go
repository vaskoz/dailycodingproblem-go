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
