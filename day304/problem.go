package day304

// ProbKnightOnBoard returns the probability that the knight is
// on the board after k-random moves from a given start position.
// Assumes an 8x8 chessboard.
// Uses O(N*N*k) extra space and O(N*N*k) in run time.
func ProbKnightOnBoard(row, col, k int) float64 {
	const N = 8
	moves := []struct{ x, y int }{
		{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2},
	}
	probs := make([][][]float64, N)
	for i := range probs {
		probs[i] = make([][]float64, N)
		for j := range probs[i] {
			probs[i][j] = make([]float64, k+1)
			probs[i][j][0] = 1.0
		}
	}
	for step := 1; step < k+1; step++ {
		for r := range probs {
			for c := range probs[r] {
				p := 0.0
				for _, move := range moves {
					nextRow := r + move.x
					nextCol := c + move.y
					if nextRow >= 0 && nextRow < 8 && nextCol >= 0 && nextCol < 8 {
						p += probs[nextRow][nextCol][step-1] / 8.0
					}
				}
				probs[r][c][step] = p
			}
		}
	}
	return probs[row][col][k]
}
