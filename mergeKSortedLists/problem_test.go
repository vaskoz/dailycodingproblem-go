package mergeklist

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input    [][]int
	expected []int
}{
	{[][]int{
		[]int{10, 15, 30},
		[]int{12, 15, 20},
		[]int{17, 20, 32}},
		[]int{10, 12, 15, 15, 17, 20, 20, 30, 32}},
	{[][]int{
		[]int{1, 2, 3},
		[]int{2, 4, 6},
		[]int{1, 2, 3},
		[]int{2, 5, 7}},
		[]int{1, 1, 2, 2, 2, 2, 3, 3, 4, 5, 6, 7}},
	{[][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
}

func TestMergeKSortedLists(t *testing.T) {
	t.Parallel()
	for i, tc := range testcases {
		if result := MergeKSortedLists(tc.input); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("did not merge list correctly for testcase#%d, got %v", i, result)
		}
	}
}

func BenchmarkMergeKSortedLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MergeKSortedLists(tc.input)
		}
	}
}

func TestMergeKSortedListsUsingHeap(t *testing.T) {
	t.Parallel()
	for i, tc := range testcases {
		if result := MergeKSortedListsUsingHeap(tc.input); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("did not merge list correctly for testcase#%d, got %v", i, result)
		}
	}
}

func BenchmarkMergeKSortedListsUsingHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MergeKSortedListsUsingHeap(tc.input)
		}
	}
}
