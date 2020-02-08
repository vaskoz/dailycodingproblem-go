package day441

import "testing"

// nolint
var testcases = []struct {
	lst   []int
	pivot int
}{
	{[]int{9, 12, 3, 5, 14, 10, 10}, 10},
	{[]int{9, 12, 3, 5, 14, 10, 10}, 20},
	{[]int{9, 12, 3, 5, 14, 10, 10}, 1},
}

func checkInvariant(lst []int, pivot int) bool {
	var i int
	for i < len(lst) && lst[i] < pivot {
		i++
	}

	for i < len(lst) && lst[i] == pivot {
		i++
	}

	for i < len(lst) && lst[i] > pivot {
		i++
	}

	return i == len(lst)
}

func TestDNFPartitionBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		copied := append([]int{}, tc.lst...)
		DNFPartitionBrute(copied, tc.pivot)

		if !checkInvariant(copied, tc.pivot) {
			t.Errorf("Didn't partition correctly. pivot=%d. Given %v got %v", tc.pivot, tc.lst, copied)
		}
	}
}

func BenchmarkDNFPartitionBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := append([]int{}, tc.lst...)
			DNFPartitionBrute(copied, tc.pivot)
		}
	}
}

func TestDNFPartition(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		copied := append([]int{}, tc.lst...)
		DNFPartition(copied, tc.pivot)

		if !checkInvariant(copied, tc.pivot) {
			t.Errorf("Didn't partition correctly. pivot=%d. Given %v got %v", tc.pivot, tc.lst, copied)
		}
	}
}

func BenchmarkDNFPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := append([]int{}, tc.lst...)
			DNFPartition(copied, tc.pivot)
		}
	}
}
