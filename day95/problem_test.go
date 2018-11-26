package day95

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	in, out []int
}{
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{1, 3, 2}, []int{2, 1, 3}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
}

func TestNextPermutation(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		in := make([]int, len(tc.in))
		copy(in, tc.in)
		NextPermutation(in)
		if !reflect.DeepEqual(in, tc.out) {
			t.Errorf("Expected %v got %v", tc.out, in)
		}
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			in := make([]int, len(tc.in))
			copy(in, tc.in)
			NextPermutation(in)
		}
	}
}
