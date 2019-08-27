package day271

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// nolint
var testcases = []struct {
	sorted   []int
	x        int
	contains bool
}{
	{[]int{1, 10, 11, 12, 14, 15}, 11, true},
	{[]int{1, 10, 11, 12, 14, 15}, 7, false},
}

func TestBinarySearch(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if index := BinarySearch(tc.sorted, tc.x); (index > 0) != tc.contains {
			t.Errorf("TCID %d, expected %v, got %v", tcid, tc.contains, (index > 0))
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BinarySearch(tc.sorted, tc.x)
		}
	}
}

func TestBinarySearchSortPkg(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if index := BinarySearchSortPkg(tc.sorted, tc.x); (index >= 0) != tc.contains {
			t.Errorf("TCID %d, expected %v, got %v", tcid, tc.contains, (index > 0))
		}
	}
}

func BenchmarkBinarySearchSortPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BinarySearchSortPkg(tc.sorted, tc.x)
		}
	}
}

func BenchmarkBinarySearchLargeInput(b *testing.B) {
	input := make([]int, 1000000)
	rand.Seed(time.Now().UnixNano())
	for i := range input {
		input[i] = rand.Int()
	}
	sort.Ints(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(input, 1000)
	}
}
