package day310

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	num     uint64
	setBits []int
}{
	{3, []int{1, 1, 2}},
}

func TestCountSetMathBits(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := CountSetMathBits(tc.num); !reflect.DeepEqual(result, tc.setBits) {
			t.Errorf("Expected %v, got %v", tc.setBits, result)
		}
	}
}

func BenchmarkCountSetMathBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CountSetMathBits(tc.num)
		}
	}
}
