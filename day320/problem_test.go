package day320

import "testing"

// nolint
var testcases = []struct {
	input        string
	windowLength int
}{
	{"jiujitsu", 5},
}

func TestSmallestWindowEveryDistinctLength(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SmallestWindowEveryDistinctLength(tc.input); result != tc.windowLength {
			t.Errorf("Expected %v, got %v", tc.windowLength, result)
		}
	}
}

func BenchmarkSmallestWindowEveryDistinctLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestWindowEveryDistinctLength(tc.input)
		}
	}
}
