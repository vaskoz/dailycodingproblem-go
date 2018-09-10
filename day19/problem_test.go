package day19

import "testing"

var testcases = []struct {
	houses   Houses
	expected int
}{
	{Houses{PaintingCosts{2, 4, 5}, PaintingCosts{3, 6, 9}}, 7},
}

func TestMinPaintingCost(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MinPaintingCost(tc.houses); result != tc.expected {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkMinPaintingCost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinPaintingCost(tc.houses)
		}
	}
}
