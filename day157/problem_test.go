package day157

import "testing"

var testcases = []struct {
	s        string
	expected bool
}{
	{"carrace", true},
	{"daily", false},
}

func TestIsPermutationPalindrome(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsPermutationPalindrome(tc.s); result != tc.expected {
			t.Errorf("For %v got %v got %v", tc.s, result, tc.expected)
		}
	}
}

func BenchmarkIsPermutationPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsPermutationPalindrome(tc.s)
		}
	}
}
