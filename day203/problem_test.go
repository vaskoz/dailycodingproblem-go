package day203

import "testing"

var testcases = []struct {
	nums     []int
	smallest int
}{
	{[]int{5, 7, 10, 3, 4}, 3},
	{[]int{3, 4, 5, 7, 10}, 3},
	{[]int{10, 3, 4, 5, 7}, 3},
	{[]int{7, 10, 3, 4, 5}, 3},
	{[]int{4, 5, 7, 10, 3}, 3},
}

func TestFindSmallestOnRotated(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if smallest := FindSmallestOnRotated(tc.nums); smallest != tc.smallest {
			t.Errorf("Expected %v, got %v", tc.smallest, smallest)
		}
	}
}

func BenchmarkFindSmallestOnRotated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindSmallestOnRotated(tc.nums)
		}
	}
}
