package day306

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	nums   []int
	k      int
	sorted []int
}{
	{[]int{3, 2, 1, 6, 5, 4, 9, 8, 7}, 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{6, 5, 3, 2, 8, 10, 9}, 3, []int{2, 3, 5, 6, 8, 9, 10}},
	{[]int{10, 9, 8, 7, 4, 70, 60, 50}, 4, []int{4, 7, 8, 9, 10, 50, 60, 70}},
}

func TestSortK(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if sorted := SortK(tc.nums, tc.k); !reflect.DeepEqual(sorted, tc.sorted) {
			t.Errorf("Expected %v, got %v", tc.sorted, sorted)
		}
	}
}

func BenchmarkSortK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SortK(tc.nums, tc.k)
		}
	}
}
