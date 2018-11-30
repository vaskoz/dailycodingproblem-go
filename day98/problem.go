package day98

// Board is a 2D grid of characters.
type Board [][]rune

// Position represents a Row and Col position
// on the 2D board.
type Position struct {
	Row, Col int
}

// DoesExistSeqAdj returns true if the word exists sequentially in
// adjacent cells. Cells can not be reused to achieve this goal.
// Returns false otherwise.
func DoesExistSeqAdj(board Board, word string) bool {
	letters := []rune(word)
	exists := false
	for row := range board {
		for col := range board[row] {
			exists = exists || search(board, row, col, letters, make(map[Position]struct{}))
			if exists {
				return true
			}
		}
	}
	return false
}

func search(board Board, row, col int, letters []rune, used map[Position]struct{}) bool {
	if len(letters) == 0 {
		return true
	}
	if board[row][col] != letters[0] {
		return false
	}
	pos := Position{row, col}
	if _, found := used[pos]; found {
		return false
	}
	used[pos] = struct{}{}
	// go up
	if search(board, up(row, len(board)), col, letters[1:], used) {
		return true
	}
	// go left
	if search(board, row, left(col, len(board[0])), letters[1:], used) {
		return true
	}
	// go right
	if search(board, row, right(col, len(board[0])), letters[1:], used) {
		return true
	}
	// go down
	if search(board, down(row, len(board)), col, letters[1:], used) {
		return true
	}

	delete(used, pos)
	return false
}

func up(row, limit int) int {
	var r int
	if row == 0 {
		r = limit - 1
	} else {
		r = row - 1
	}
	return r
}

func down(row, limit int) int {
	var r int
	if row == limit-1 {
		r = 0
	} else {
		r = row + 1
	}
	return r
}

func left(col, limit int) int {
	var c int
	if col == 0 {
		c = limit - 1
	} else {
		c = col - 1
	}
	return c
}

func right(col, limit int) int {
	var c int
	if col == limit-1 {
		c = 0
	} else {
		c = col + 1
	}
	return c
}
