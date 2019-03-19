package day205

import (
	"testing"
)

var testcases = []struct {
	in  uint
	out uint
	err error
}{
	{123, 132, nil},
	{132, 213, nil},
	{321, 0, ErrNoGreaterPermutation()},
}

func TestNextPermutationUint(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := NextPermutationUint(tc.in); result != tc.out || err != tc.err {
			t.Errorf("Expected (%v,%v) got (%v,%v)", tc.out, tc.err, result, err)
		}
	}
}

func BenchmarkNextPermutationUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NextPermutationUint(tc.in)
		}
	}
}
