package day143

import "testing"

var testcases = []struct {
	lst   []int
	pivot int
}{
	{[]int{9, 12, 3, 5, 14, 10, 10}, 10},
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

func TestPartitionBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		copied := append([]int{}, tc.lst...)
		PartitionBrute(copied, tc.pivot)
		if !checkInvariant(copied, tc.pivot) {
			t.Errorf("Didn't partition correctly. pivot=%d. Given %v got %v", tc.pivot, tc.lst, copied)
		}
	}
}

func BenchmarkPartitionBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := append([]int{}, tc.lst...)
			PartitionBrute(copied, tc.pivot)
		}
	}
}

func TestPartition(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		copied := append([]int{}, tc.lst...)
		Partition(copied, tc.pivot)
		if !checkInvariant(copied, tc.pivot) {
			t.Errorf("Didn't partition correctly. pivot=%d. Given %v got %v", tc.pivot, tc.lst, copied)
		}
	}
}

func BenchmarkPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := append([]int{}, tc.lst...)
			Partition(copied, tc.pivot)
		}
	}
}
