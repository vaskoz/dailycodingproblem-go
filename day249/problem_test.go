package day249

import "testing"

// nolint
var testcases = []struct {
	nums   []int64
	maxXor int64
}{
	{[]int64{30, -20, 10, 100, 3300, 1001}, 3853},
	{[]int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 15},
}

func TestMaxXorPairSort(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxXorPairSort(tc.nums); result != tc.maxXor {
			t.Errorf("Expected %v, got %v", tc.maxXor, result)
		}
	}
}

func BenchmarkMaxXorPairSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxXorPairSort(tc.nums)
		}
	}
}

func TestMaxXorPairBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxXorPairBrute(tc.nums); result != tc.maxXor {
			t.Errorf("Expected %v, got %v", tc.maxXor, result)
		}
	}
}

func BenchmarkMaxXorPairBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxXorPairBrute(tc.nums)
		}
	}
}
