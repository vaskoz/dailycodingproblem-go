package day27

import "testing"

var testcases = []struct {
	input    string
	expected bool
}{
	{"([])[]({})", true},
	{"([)]", false},
	{"((()", false},
	{"((())))", false},
}

func TestWellBalanced(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := WellBalanced(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v for %v", tc.expected, result, tc.input)
		}
	}
}

func BenchmarkWellBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			WellBalanced(tc.input)
		}
	}
}
