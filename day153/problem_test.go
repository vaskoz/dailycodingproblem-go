package day153

import "testing"

var testcases = []struct {
	text, one, two string
	expected       int
}{
	{"dog cat hello cat dog dog hello cat world", "hello", "world", 1},
	{"dog cat hello cat dog dog hello cat world", "hello", "cat", 0},
	{"dog cat hello cat dog dog hello cat world", "wat?", "cat", -1},
}

func TestSmallestDistanceBetweenWords(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SmallestDistanceBetweenWords(tc.text, tc.one, tc.two); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSmallestDistanceBetweenWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestDistanceBetweenWords(tc.text, tc.one, tc.two)
		}
	}
}
