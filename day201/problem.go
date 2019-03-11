package day201

// MaxWeightPathTriangle returns the weight of the maximum weight path
// as defined by the problem.
// Runs in O(N) time.
// Assumes a correct triangle as input.
func MaxWeightPathTriangle(tri [][]int) int {
	return maxWeightPathTriangleHelper(tri, 0, 0)
}

func maxWeightPathTriangleHelper(tri [][]int, row, col int) int {
	if row == len(tri)-1 {
		return tri[row][col]
	}
	left := maxWeightPathTriangleHelper(tri, row+1, col)
	right := maxWeightPathTriangleHelper(tri, row+1, col+1)
	return tri[row][col] + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
