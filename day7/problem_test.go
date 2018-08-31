package day7

import "testing"

var testcases = []struct {
	input    string
	expected int
}{
	{"111", 3},
	{"1111", 5},
	{"2626", 4},
	{"1234", 3},
}

func TestNumberOfDecodings(t *testing.T) {
	tc.Parallel()
	for _, tc := range testcases {
		if result := NumberOfDecodings(tc.input); result != tc.expected {
			t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, result)
		}
	}
}

func BenchmarkNumberOfDecodings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NumberOfDecodings(tc.input)
		}
	}
}
