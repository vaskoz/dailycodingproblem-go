package day468

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	matrix, rotated Matrix
}{
	{Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, Matrix{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}},
}

func TestRotateMatrixRightClockwise90(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := RotateMatrixRightClockwise90(tc.matrix); !reflect.DeepEqual(result, tc.rotated) {
			t.Errorf("Expected %v got %v", tc.rotated, result)
		}
	}
}

func BenchmarkRotateMatrixRightClockwise90(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RotateMatrixRightClockwise90(tc.matrix)
		}
	}
}
