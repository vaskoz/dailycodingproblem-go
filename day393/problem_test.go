package day393

import "testing"

// nolint
var testcases = []struct {
	nums       []int
	start, end int
}{
	{[]int{9, 6, 1, 3, 8, 10, 12, 11}, 8, 12},
	{[]int{}, 0, 0},
	{nil, 0, 0},
	{[]int{12}, 12, 12},
	{[]int{9, 1, 3, 8, 10, 12, 11, 2, 4, 5}, 1, 5},
}

func TestLargestRangeBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if s, e := LargestRangeBrute(tc.nums); s != tc.start || e != tc.end {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.start, tc.end, s, e)
		}
	}
}

func BenchmarkLargestRangeBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LargestRangeBrute(tc.nums)
		}
	}
}

func TestLargestRangeSort(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if s, e := LargestRangeSort(tc.nums); s != tc.start || e != tc.end {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.start, tc.end, s, e)
		}
	}
}

func BenchmarkLargestRangeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LargestRangeSort(tc.nums)
		}
	}
}
