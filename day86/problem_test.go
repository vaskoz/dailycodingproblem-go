package day86

import "testing"

var testcases = []struct {
	input    string
	expected int
}{
	{"()())()", 1},
	{")(", 2},
	{"((()(())))", 0},
	{"(((()(())))", 1},
	{"((()(()))))", 1},
}

func TestMinRemovedParens(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinRemovedParens(tc.input); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkMinRemovedParens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinRemovedParens(tc.input)
		}
	}
}
