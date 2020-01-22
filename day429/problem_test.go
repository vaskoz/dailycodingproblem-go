package day429

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	row      int
	expected []int
}{
	{0, []int{1}},
	{1, []int{1, 1}},
	{2, []int{1, 2, 1}},
	{3, []int{1, 3, 3, 1}},
	{4, []int{1, 4, 6, 4, 1}},
	{5, []int{1, 5, 10, 10, 5, 1}},
	{6, []int{1, 6, 15, 20, 15, 6, 1}},
	{7, []int{1, 7, 21, 35, 35, 21, 7, 1}},
}

func TestPascalTriangleRow(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := PascalTriangleRow(tc.row); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkPascalTriangleRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PascalTriangleRow(tc.row)
		}
	}
}
