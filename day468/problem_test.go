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
	{Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}, Matrix{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}},
	{Matrix{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}, Matrix{
		{21, 16, 11, 6, 1},
		{22, 17, 12, 7, 2},
		{23, 18, 13, 8, 3},
		{24, 19, 14, 9, 4},
		{25, 20, 15, 10, 5},
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

func TestRotateMatrixRightClockwise90InPlace(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		copied := copyMatrix(tc.matrix)
		RotateMatrixRightClockwise90InPlace(copied)

		if !reflect.DeepEqual(copied, tc.rotated) {
			t.Errorf("Expected %v got %v", tc.rotated, copied)
		}
	}
}

func BenchmarkRotateMatrixRightClockwise90InPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := copyMatrix(tc.matrix)
			RotateMatrixRightClockwise90InPlace(copied)
		}
	}
}

func copyMatrix(m Matrix) Matrix {
	copied := make(Matrix, len(m))

	for i := range m {
		copied[i] = make([]int, len(m))
		copy(copied[i], m[i])
	}

	return copied
}
