package day472

import "testing"

// nolint
var testcases = []struct {
	input    string
	expected int
}{
	{"111", 3},
	{"1111", 5},
	{"2626", 4},
	{"1234", 3},
	{"1111127", 13},
}

func TestNumberOfDecodings(t *testing.T) {
	t.Parallel()

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
