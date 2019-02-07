package day168

// Matrix is a 2D array of integers of NxN dimension.
type Matrix [][]int

// RotateMatrixRightClockwise90 does what it says.
// It creates a new Matrix for the result.
func RotateMatrixRightClockwise90(m Matrix) Matrix {
	result := make(Matrix, len(m))
	for i := range result {
		result[i] = make([]int, len(m))
	}
	for i := range m {
		for j := range m[i] {
			result[j][len(m)-1-i] = m[i][j]
		}
	}
	return result
}
