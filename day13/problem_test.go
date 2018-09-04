package day13

import "testing"

var testcases = []struct {
	input    string
	k        int
	expected string
}{
	{"abcba", 2, "bcb"},
	{"aa", 2, "aa"},
	{"xyzabababbbbaaaac", 2, "abababbbbaaaa"},
	{"xyzabababbbbaaaac", 2, "abababbbbaaaa"},
}

func TestLongestSubKDistinct(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LongestSubKDistinct(tc.input, tc.k); result != tc.expected {
			t.Errorf("Given %v I expected %v but got %v with k %v", tc.input, tc.expected, result, tc.k)
		}
	}
}

func BenchmarkLongestSubKDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestSubKDistinct(tc.input, tc.k)
		}
	}
}
