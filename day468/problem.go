package day468

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

// RotateMatrixRightClockwise90InPlace does what it says.
func RotateMatrixRightClockwise90InPlace(m Matrix) {
	n := len(m)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			var next int
			next, m[j][n-i-1] = m[j][n-i-1], m[i][j]      // right-side
			next, m[n-i-1][n-j-1] = m[n-i-1][n-j-1], next // bottom-side
			next, m[n-j-1][i] = m[n-j-1][i], next         // left-side
			m[i][j] = next                                // top-side
		}
	}
}
