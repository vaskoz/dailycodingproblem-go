package day195

// Matrix is a M by N matrix of integers.
type Matrix [][]int

// CountLargerAndSmallerBrute computes the number of elements of M smaller
// than M[i1, j1] and larger than M[i2, j2].
func CountLargerAndSmallerBrute(mat Matrix, i1, j1, i2, j2 int) int {
	count := 0
	smallerThan := mat[i1][j1]
	largerThan := mat[i2][j2]
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] < smallerThan {
				count++
			} else if mat[i][j] > largerThan {
				count++
			}
		}
	}
	return count
}
