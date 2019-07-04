package day315

// IsToeplitzMatrix answers if a given matrix is a Toeplitz Matrix.
// https://en.wikipedia.org/wiki/Toeplitz_matrix
// Runs in O(N) time and O(1) extra space.
func IsToeplitzMatrix(mat [][]int) bool {
	for row := 0; row < len(mat)-1; row++ {
		r := row
		c := 0
		num := mat[r][c]
		for r < len(mat) && c < len(mat[r]) {
			if mat[r][c] != num {
				return false
			}
			r++
			c++
		}
	}
	for col := 1; col < len(mat[0])-1; col++ {
		r := 0
		c := col
		num := mat[r][c]
		for r < len(mat) && c < len(mat[r]) {
			if mat[r][c] != num {
				return false
			}
			r++
			c++
		}
	}
	return true
}
