package day58

import "testing"

var testcases = []struct {
	sorted   []int
	target   int
	expected int
}{
	{[]int{13, 18, 25, 2, 8, 10}, 8, 4},
	{[]int{13, 18, 25, 27, 2, 8, 10}, 8, 5},
	{[]int{13, 18, 25, 27, 2, 8, 10}, 18, 1},
	{[]int{13, 18, 25, 27, 2, 8, 10}, 14, -1},
	{[]int{13, 18, 25, 27, 30, 36, 54, 100, 2}, 2, 8},
}

func TestIndexSortedRotated(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if index := IndexSortedRotated(tc.sorted, tc.target); index != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, index)
		}
	}
}

func BenchmarkIndexSortedRotated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IndexSortedRotated(tc.sorted, tc.target)
		}
	}
}
