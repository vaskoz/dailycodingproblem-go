package sudoku

// Board is a 1-D array of size n^2 x n^2. You need 'n' to know how to read it properly.
type Board []int

// Solver returns a solved Sudoku puzzle of size n^2 x n^2.
func Solver(board Board, n int) (Board, bool) {
	row, col, complete := findEmpty(board, n)
	if complete {
		return board, complete
	}
	for choice := 1; choice <= n*n; choice++ {
		isValid := true
		gridrow, gridcol := row/n, col/n
		for r := n * gridrow; r < n*gridrow+n; r++ {
			for c := n * gridcol; c < n*gridcol+n; c++ {
				if board[r*n*n+c] == choice {
					isValid = false
				}
			}
		}
		for j := 0; j < n*n; j++ {
			if board[n*n*j+col] == choice || board[row*n*n+j] == choice {
				isValid = false
				break
			}
		}
		if isValid {
			board[row*n*n+col] = choice
			var solved bool
			board, solved = Solver(board, n)
			if solved {
				return board, solved
			}
			board[row*n*n+col] = 0
		}
	}
	return board, false
}

func findEmpty(board Board, n int) (int, int, bool) {
	for i, box := range board {
		if box == 0 {
			return i / (n * n), i % (n * n), false
		}
	}
	return 0, 0, true
}
