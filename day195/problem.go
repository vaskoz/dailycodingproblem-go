package day195

// Matrix is a M by N matrix of integers.
type Matrix [][]int

// CountLargerAndSmallerBrute computes the number of elements of M smaller
// than M[i1, j1] and larger than M[i2, j2].
// Runs in O(M*N) time. Scans the entire matrix.
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

// CountLargerAndSmallerEfficient computes the number of elements of M smaller
// than M[i1, j1] and larger than M[i2, j2].
// Runs in O(Answer). Only examines the number of elements in the answer.
func CountLargerAndSmallerEfficient(mat Matrix, i1, j1, i2, j2 int) int {
	count := 0
	smallerThan := mat[i1][j1]
	for i := range mat {
		var found bool
		for _, v := range mat[i] {
			if v < smallerThan {
				found = true
				count++
			} else {
				break
			}
		}
		if !found {
			break
		}
	}
	largerThan := mat[i2][j2]
	for i := range mat {
		var found bool
		for j := range mat[len(mat)-1-i] {
			if mat[len(mat)-1-i][len(mat[i])-1-j] > largerThan {
				found = true
				count++
			} else {
				break
			}
		}
		if !found {
			break
		}
	}
	return count
}
