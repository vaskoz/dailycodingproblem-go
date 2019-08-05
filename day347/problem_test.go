package day347

import "testing"

// nolint
var testcases = []struct {
	str      string
	k        int
	expected string
}{
	{"daily", 1, "ailyd"},
	{"gaurang", 3, "agangru"},
	{"geeksforgeeks", 5, "eefggeekkorss"},
}

func TestLexicographicallySmallest(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LexicographicallySmallest(tc.str, tc.k); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkLexicographicallySmallest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LexicographicallySmallest(tc.str, tc.k)
		}
	}
}
