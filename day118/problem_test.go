package day118

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input, expected []int
}{
	{[]int{-9, -2, 0, 2, 3}, []int{0, 4, 4, 9, 81}},
	{[]int{-20, -19, -18, -17, -16, -15, -14, -13, -12, -11, -10, -9,
		-2, 0, 2, 3, 4, 5, 7, 10, 12, 14, 18, 20, 22, 24, 25, 26, 27, 29, 30},
		[]int{0, 4, 4, 9, 16, 25, 49, 81, 100, 100, 121, 144, 144, 169, 196, 196,
			225, 256, 289, 324, 324, 361, 400, 400, 484, 576, 625, 676, 729, 841, 900}},
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

func TestSquareElementsSortedLinear(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SquareElementsSortedLinear(tc.input); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}
func BenchmarkSquareElementsSortedLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SquareElementsSortedLinear(tc.input)
		}
	}
}
