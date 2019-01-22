package day149

import "testing"

var testcases = []struct {
	L          []int
	start, end int
	expected   int
}{
	{[]int{1, 2, 3, 4, 5}, 1, 3, 5},
}

func TestSumSublist(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SumSublist(tc.L, tc.start, tc.end); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSumSublist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SumSublist(tc.L, tc.start, tc.end)
		}
	}
}
