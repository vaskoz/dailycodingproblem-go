package day201

import "testing"

var testcases = []struct {
	triangle  [][]int
	maxWeight int
}{
	{[][]int{{1}, {2, 3}, {1, 5, 1}}, 9},
}

func TestMaxWeightPathTriangle(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxWeightPathTriangle(tc.triangle); result != tc.maxWeight {
			t.Errorf("Expected %v, got %v", tc.maxWeight, result)
		}
	}
}

func BenchmarkMaxWeightPathTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxWeightPathTriangle(tc.triangle)
		}
	}
}
