package day113

import "testing"

var testcases = []struct {
	input, expected string
}{
	{"hello world here", "here world hello"},
}

func TestReverseWords(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ReverseWords(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ReverseWords(tc.input)
		}
	}
}

func TestReverseWordsInPlace(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ReverseWordsInPlace(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkReverseWordsInPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ReverseWordsInPlace(tc.input)
		}
	}
}
