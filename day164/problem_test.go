package day164

import "testing"

var testcases = []struct {
	nums []int
	dup  int
}{
	{[]int{3, 1, 5, 4, 2, 4}, 4},
}

func TestFindDuplicate(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if dup := FindDuplicate(tc.nums); dup != tc.dup {
			t.Errorf("Expected %v got %v", tc.dup, dup)
		}
	}
}

func BenchmarkFindDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindDuplicate(tc.nums)
		}
	}
}
