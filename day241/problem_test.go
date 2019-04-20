package day241

import "testing"

var testcases = []struct {
	citations []int
	hindex    int
}{
	{[]int{4, 3, 0, 1, 5}, 3},
	{[]int{10}, 1},
}

func TestHIndex(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := HIndex(tc.citations); result != tc.hindex {
			t.Errorf("Expected %v, got %v", tc.hindex, result)
		}
	}
}

func BenchmarkHIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			HIndex(tc.citations)
		}
	}
}
