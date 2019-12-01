package day377

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input    []float64
	k        int
	expected []float64
}{
	{
		[]float64{-1, 5, 13, 8, 2, 3, 3, 1},
		3,
		[]float64{5, 8, 8, 3, 3, 3},
	},
	{
		[]float64{-1, 5, 13, 8, 2, 3, 3, 1},
		4,
		[]float64{6.5, 6.5, 5.5, 3, 2.5},
	},
}

func TestBruteKMedians(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := BruteKMedians(tc.input, tc.k); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkBruteKMedians(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BruteKMedians(tc.input, tc.k)
		}
	}
}
