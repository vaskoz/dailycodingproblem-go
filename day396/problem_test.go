package day396

import "testing"

// nolint
var testcases = []struct {
	str      string
	expected int
}{
	{"MAPTPTMTPA", 7},
	{"ABFOOBARRABOOFYZ", 12},
}

func TestLongestPalindromicSubsequenceBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := LongestPalindromicSubsequenceBrute(tc.str); res != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkLongestPalindromicSubsequenceBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestPalindromicSubsequenceBrute(tc.str)
		}
	}
}

func TestLongestPalindromicSubsequenceDP(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := LongestPalindromicSubsequenceDP(tc.str); res != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkLongestPalindromicSubsequenceDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestPalindromicSubsequenceDP(tc.str)
		}
	}
}
