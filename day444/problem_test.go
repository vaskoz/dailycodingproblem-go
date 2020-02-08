package day444

import "testing"

// nolint
var testcases = []struct {
	input, substr string
	expected      int
}{
	{"hello world!", "wor", 6},
	{"hello world!", "buddy", -1},
	{"hello world!", "hello", 0},
	{"hello world!", "ll", 2},
	{"hello world!", "llo w", 2},
	{"hello world!", "ld", 9},
}

func TestStringMatchKMP(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := StringMatchKMP(tc.input, tc.substr); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkStringMatchKMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			StringMatchKMP(tc.input, tc.substr)
		}
	}
}
