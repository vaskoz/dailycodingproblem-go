package day123

import "testing"

var testcases = []struct {
	input    string
	isNumber bool
}{
	{"10", true},
	{"-10", true},
	{"10.1", true},
	{"-10.1", true},
	{"1e5", true},
	{"a", false},
	{"x 1", false},
	{"a -2", false},
	{"-", false},
	{"1-2", false},
	{".12", false},
	{"0.12", true},
	{"1eE2", false},
	{"1E2", true},
}

func TestIsANumber(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsANumber(tc.input); result != tc.isNumber {
			t.Errorf("For input %v expected %v got %v", tc.input, tc.isNumber, result)
		}
	}
}

func BenchmarkIsANumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsANumber(tc.input)
		}
	}
}
