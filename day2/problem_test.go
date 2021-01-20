package day2

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input, expected []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	{[]int{3, 2, 1}, []int{2, 3, 6}},
}

func TestProductsExceptI(t *testing.T) {
	t.Run("ProductsExceptI", func(t *testing.T) {
		t.Parallel()
		for _, tc := range testcases {
			if result := ProductsExceptI(tc.input); !reflect.DeepEqual(result, tc.expected) {
				t.Error("didn't get what I expected")
			}
		}
	})
	t.Run("ProductsExceptINoDivisionLinear", func(t *testing.T) {
		t.Parallel()
		for _, tc := range testcases {
			if result := ProductsExceptINoDivisionLinear(tc.input); !reflect.DeepEqual(result, tc.expected) {
				t.Error("didn't get what I expected")
			}
		}
	})
	t.Run("ProductsExceptINoDivision", func(t *testing.T) {
		t.Parallel()
		for _, tc := range testcases {
			if result := ProductsExceptINoDivision(tc.input); !reflect.DeepEqual(result, tc.expected) {
				t.Error("didn't get what I expected")
			}
		}
	})
}

func BenchmarkProductsExceptI(b *testing.B) {
	b.Run("ProductsExceptI", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testcases {
				ProductsExceptI(tc.input)
			}
		}
	})
	b.Run("ProductsExceptINoDivisionLinear", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testcases {
				ProductsExceptINoDivisionLinear(tc.input)
			}
		}
	})
	b.Run("ProductsExceptINoDivision", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testcases {
				ProductsExceptI(tc.input)
			}
		}
	})
}
