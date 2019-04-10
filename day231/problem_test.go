package day231

import "testing"

var testcases = []struct {
	s, expected string
}{
	{"bbbaa", "babab"},
	{"aaabbc", "abcaba"},
	{"aaab", ""},
}

func TestNoAdjacent(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := NoAdjacent(tc.s); result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}

func BenchmarkNoAdjacent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NoAdjacent(tc.s)
		}
	}
}
