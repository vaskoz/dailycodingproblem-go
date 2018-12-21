package day118

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input, expected []int
}{
	{[]int{-9, -2, 0, 2, 3}, []int{0, 4, 4, 9, 81}},
}

func TestSquareElementsSortedBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SquareElementsSortedBrute(tc.input); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}
func BenchmarkSquareElementsSortedBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SquareElementsSortedBrute(tc.input)
		}
	}
}
