package day411

import "testing"

// nolint
var testcases = []struct {
	input    string
	k        int
	expected bool
}{
	{"waterrfetawx", 2, true},
}

func TestIsKPalindromeBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := IsKPalindromeBrute(tc.input, tc.k); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkIsKPalindromeBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsKPalindromeBrute(tc.input, tc.k)
		}
	}
}

func TestIsKPalindromeDP(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := IsKPalindromeDP(tc.input, tc.k); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkIsKPalindromeDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsKPalindromeDP(tc.input, tc.k)
		}
	}
}
