package day332

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	m, n     int
	expected []Pair
}{
	{
		100, 4,
		[]Pair{{48, 52}},
	},
}

func TestPositiveIntegerPairs(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if pairs := PositiveIntegerPairs(tc.m, tc.n); !reflect.DeepEqual(pairs, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, pairs)
		}
	}
}

func BenchmarkPositiveIntegerPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PositiveIntegerPairs(tc.m, tc.n)
		}
	}
}
