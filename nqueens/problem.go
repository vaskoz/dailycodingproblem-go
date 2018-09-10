package nqueens

// NQueens returns the number of possible placements of queens on a
// size x size chessboard.
func NQueens(size int, board []int) int {
	if size == len(board) {
		return 1
	}
	var count int
	for i := 0; i < size; i++ {
		board = append(board, i)
		if isValid(board) {
			count += NQueens(size, board)
		}
		board = board[:len(board)-1]
	}
	return count
}

func isValid(board []int) bool {
	curQueenR, curQueenC := len(board)-1, board[len(board)-1]
	for row := 0; row < len(board)-1; row++ {
		col := board[row]
		diff := curQueenC - col
		if diff < 0 {
			diff = -diff
		}
		if diff == 0 || diff == curQueenR-row {
			return false
		}
	}
	return true
}
