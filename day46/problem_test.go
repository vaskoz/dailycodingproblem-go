package day46

import "testing"

var testcases = []struct {
	input, expected string
}{
	{"aabcdcb", "bcdcb"},
	{"bananas", "anana"},
	{"abcdef", "a"},
	{"dfqwoef2jfeoijweconocfnmfjadsflkjqpwefnvpxczxqm", "conoc"},
	{"", ""},
}

func TestLongestPalindromicSubstringBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LongestPalindromicSubstringBrute(tc.input); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkLongestPalindromicSubstringBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestPalindromicSubstringBrute(tc.input)
		}
	}
}

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LongestPalindromicSubstring(tc.input); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkLongestPalindromicSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestPalindromicSubstring(tc.input)
		}
	}
}
