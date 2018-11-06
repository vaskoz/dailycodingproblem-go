package day76

// MinRemoveColumns determines the minimum number of columns
// that can be removed to ensure that each row is ordered
// from top to bottom lexicographically.
// Runs in O(M*N) checking all input values.
func MinRemoveColumns(matrix [][]rune) int {
	if len(matrix) < 2 {
		return 0
	}
	removals := 0
	for col := 0; col < len(matrix[0]); col++ {
		for row := 1; row < len(matrix); row++ {
			if matrix[row][col] < matrix[row-1][col] {
				removals++
				break
			}
		}
	}
	return removals
}
