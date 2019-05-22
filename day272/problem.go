package day272

// PossibleWaysToThrow returns the number of possible distinct throws
// of N dice with a given number of faces to equal a desired total.
// Runs in O(dice*faces*total) with memory O(dice*total).
func PossibleWaysToThrow(ndice, faces, total int) int {
	//# The main function that returns number of ways to get sum 'x'
	//# with 'n' dice and 'm' with m faces.
	table := make([][]int, ndice+1)
	for i := range table {
		table[i] = make([]int, total+1)
	}
	for i := 0; i < min(faces+1, total+1); i++ {
		table[1][i] = 1
	}
	for i := 2; i < ndice+1; i++ {
		for j := 1; j < total+1; j++ {
			for k := 1; k < min(faces+1, j); k++ {
				table[i][j] += table[i-1][j-k]
			}
		}
	}
	lastRow := len(table) - 1
	lastCol := len(table[lastRow]) - 1
	return table[lastRow][lastCol]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
