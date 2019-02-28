package day189

import "testing"

var testcases = []struct {
	nums   []int
	length int
}{
	{[]int{5, 1, 3, 5, 2, 3, 4, 1}, 5},
}

func TestLengthLongestDistinct(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LengthLongestDistinct(tc.nums); result != tc.length {
			t.Errorf("Expected %v, got %v", tc.length, result)
		}
	}
}

func BenchmarkLengthLongestDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LengthLongestDistinct(tc.nums)
		}
	}
}
