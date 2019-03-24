package day214

import "testing"

var testcases = []struct {
	n                      int
	longestConsecutiveOnes int
}{
	{156, 3},
	{20432348234, 3},
	{65535, 16},
}

func TestLengthLongestConsecutiveRunOnesBitShift(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LengthLongestConsecutiveRunOnesBitShift(tc.n); result != tc.longestConsecutiveOnes {
			t.Errorf("Given n=%d, expected %d, got %d", tc.n, tc.longestConsecutiveOnes, result)
		}
	}
}

func BenchmarkLengthLongestConsecutiveRunOnesBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LengthLongestConsecutiveRunOnesBitShift(tc.n)
		}
	}
}

func TestLengthLongestConsecutiveRunOnesString(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LengthLongestConsecutiveRunOnesString(tc.n); result != tc.longestConsecutiveOnes {
			t.Errorf("Given n=%d, expected %d, got %d", tc.n, tc.longestConsecutiveOnes, result)
		}
	}
}

func BenchmarkLengthLongestConsecutiveRunOnesString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LengthLongestConsecutiveRunOnesString(tc.n)
		}
	}
}
