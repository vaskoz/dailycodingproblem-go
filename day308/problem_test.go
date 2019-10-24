package day308

import "testing"

// nolint
var testcases = []struct {
	input    []rune
	expected int
}{
	{[]rune{'F', '|', 'T', '&', 'T'}, 2},
	{[]rune{'T', '|', 'T', '&', 'F', '^', 'T'}, 4},
}

func TestCountParensEqualTrue(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := CountParensEqualTrue(tc.input); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkCountParensEqualTrue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountParensEqualTrue(tc.input)
		}
	}
}
