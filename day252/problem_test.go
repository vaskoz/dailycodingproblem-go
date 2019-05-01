package day252

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	numerator, denominator int
	denominators           []int
}{
	{4, 13, []int{4, 18, 468}},
}

func TestEgyptianFractions(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if denoms := EgyptianFractions(tc.numerator, tc.denominator); !reflect.DeepEqual(tc.denominators, denoms) {
			t.Errorf("Expected %v, got %v", tc.denominators, denoms)
		}
	}
}

func BenchmarkEgyptianFractions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EgyptianFractions(tc.numerator, tc.denominator)
		}
	}
}
