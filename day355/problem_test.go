package day355

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	x []float64
	y []int
}{
	{
		[]float64{1.3, 2.3, 4.4},
		[]int{1, 2, 5},
	},
	{
		[]float64{4.4, 1.3, 2.3},
		[]int{5, 1, 2},
	},
	{
		[]float64{1.5, 2.9, 4.55},
		[]int{1, 3, 5},
	},
}

func TestFindY(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if y := FindY(tc.x); !reflect.DeepEqual(y, tc.y) {
			t.Errorf("Expected %v, got %v", tc.y, y)
		}
	}
}

func BenchmarkFindY(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindY(tc.x)
		}
	}
}
