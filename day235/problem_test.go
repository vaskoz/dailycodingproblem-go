package day235

import "testing"

var testcases = []struct {
	nums     []int
	min, max int
}{
	{[]int{3, 3, 2, 3, 5, 6, 3, 2, 3, 3, 3, 2, -10, 323, 3, 3, 3, 35, 5, 6, 6, 6, 500, 23, 320, 320, 10}, -10, 500},
	{nil, 0, 0},
	{[]int{4}, 4, 4},
}

func TestMaxMinNotGoodEnough(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if min, max := MaxMinNotGoodEnough(tc.nums); min != tc.min || max != tc.max {
			t.Errorf("Expected (%d,%d) got (%d,%d)", tc.min, tc.max, min, max)
		}
	}
}

func BenchmarkMaxMinNotGoodEnough(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxMinNotGoodEnough(tc.nums)
		}
	}
}

func TestMaxMinPairs(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if min, max := MaxMinPairs(tc.nums); min != tc.min || max != tc.max {
			t.Errorf("Expected (%d,%d) got (%d,%d)", tc.min, tc.max, min, max)
		}
	}
}

func BenchmarkMaxMinPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxMinPairs(tc.nums)
		}
	}
}
