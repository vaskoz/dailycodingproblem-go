package day14

import (
	"math"
	"testing"
)

var testcases = []struct {
	iterations int
	expected   float64
	delta      float64
}{
	{100000, 3.1415, 0.1},
	{10000, 3.1415, 0.1},
	{1000, 3.1415, 0.1},
}

func TestEstimatePi(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if pi := EstimatePi(tc.iterations); math.Abs(pi-tc.expected) > tc.delta {
			t.Errorf("Expected %v but got %v with delta %v", tc.expected, pi, tc.delta)
		}
	}
}

func BenchmarkEstimatePi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EstimatePi(tc.iterations)
		}
	}
}
