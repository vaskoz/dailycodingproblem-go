package day64

// KnightsTourCountBacktracking returns the number of tours
// on a NxN chessboard. The only argument is N.
// Operates in exponential time. O(8^(N^2)).
func KnightsTourCountBacktracking(n int) int {
	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board := make([][]bool, n)
			for x := range board {
				board[x] = make([]bool, n)
			}
			board[i][j] = true
			count += knightsTourCounter(n, board, [][]int{{i, j}})
		}
	}
	return count
}

func knightsTourCounter(n int, board [][]bool, tour [][]int) int {
	if len(tour) == n*n {
		return 1
	}
	var count int
	last := tour[len(tour)-1]
	r, c := last[0], last[1]
	for _, move := range validMoves(board, r, c, n) {
		board[move[0]][move[1]] = true
		count += knightsTourCounter(n, board, append(tour, move))
		board[move[0]][move[1]] = false
	}
	return count
}

func validMoves(board [][]bool, r, c, n int) [][]int {
	moves := make([][]int, 0, 8)
	deltas := [][]int{{2, 1}, {1, 2}, {1, -2}, {-2, 1}, {-1, 2}, {2, -1}, {-1, -2}, {-2, -1}}
	for _, delta := range deltas {
		if isValidMove(board, n, r+delta[0], c+delta[1]) {
			moves = append(moves, []int{r + delta[0], c + delta[1]})
		}
	}
	return moves
}

func isValidMove(board [][]bool, n, r, c int) bool {
	return 0 <= r && r < n && 0 <= c && c < n && !board[r][c]
}
