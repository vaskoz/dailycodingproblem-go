package day425

// Chessboard is a slice of strings. The first string represents the top row
// and the last string represents the bottom row of the board.
// The top is the black player's side and the bottom is the white player's side.
type Chessboard []string

// IsSoloBlackKingInCheck answers if a solo black king is in check.
// All other pieces on the board are white pieces.
func IsSoloBlackKingInCheck(board Chessboard) bool {
	kR, kC := findBlackKing(board)

	for r := range board {
		for c := range board[r] {
			switch board[r][c] {
			case 'B':
				if bishop(board, r, c, kR, kC) {
					return true
				}
			case 'N':
				if knight(board, r, c, kR, kC) {
					return true
				}
			case 'P':
				if pawn(board, r, c, kR, kC) {
					return true
				}
			case 'R':
				if rook(board, r, c, kR, kC) {
					return true
				}
			case 'Q':
				if queen(board, r, c, kR, kC) {
					return true
				}
			}
		}
	}

	return false
}

func knight(_ Chessboard, nR, nC, kR, kC int) bool {
	return (nR-2 == kR && nC+1 == kC) ||
		(nR-1 == kR && nC+2 == kC) ||
		(nR+1 == kR && nC+2 == kC) ||
		(nR+2 == kR && nC+1 == kC) ||
		(nR+2 == kR && nC-1 == kC) ||
		(nR+1 == kR && nC-2 == kC) ||
		(nR-1 == kR && nC-2 == kC) ||
		(nR-2 == kR && nC-1 == kC)
}

func pawn(_ Chessboard, pR, pC, kR, kC int) bool {
	return (pR-1 == kR && pC-1 == kC) || (pR-1 == kR && pC+1 == kC)
}

func queen(board Chessboard, qR, qC, kR, kC int) bool {
	return rook(board, qR, qC, kR, kC) || bishop(board, qR, qC, kR, kC)
}

func rook(board Chessboard, rR, rC, kR, kC int) bool {
	// up
	for r := rR - 1; r >= 0; r-- {
		if r == kR && rC == kC {
			return true
		} else if board[r][rC] != '.' {
			break
		}
	}
	// left
	for c := rC - 1; c >= 0; c-- {
		if rR == kR && c == kC {
			return true
		} else if board[rR][c] != '.' {
			break
		}
	}
	// right
	for c := rC + 1; c < 8; c++ {
		if rR == kR && c == kC {
			return true
		} else if board[rR][c] != '.' {
			break
		}
	}
	// down
	for r := rR + 1; r < 8; r++ {
		if r == kR && rC == kC {
			return true
		} else if board[r][rC] != '.' {
			break
		}
	}

	return false
}

func bishop(board Chessboard, bR, bC, kR, kC int) bool {
	// up-right
	for r, c := bR-1, bC+1; r >= 0 && c < 8; r, c = r-1, c+1 {
		if r == kR && c == kC {
			return true
		} else if board[r][c] != '.' {
			break
		}
	}
	// up-left
	for r, c := bR-1, bC-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if r == kR && c == kC {
			return true
		} else if board[r][c] != '.' {
			break
		}
	}
	// down-right
	for r, c := bR+1, bC+1; r < 8 && c < 8; r, c = r+1, c+1 {
		if r == kR && c == kC {
			return true
		} else if board[r][c] != '.' {
			break
		}
	}
	// down-left
	for r, c := bR+1, bC-1; r < 8 && c >= 0; r, c = r+1, c-1 {
		if r == kR && c == kC {
			return true
		} else if board[r][c] != '.' {
			break
		}
	}

	return false
}

func findBlackKing(board Chessboard) (row int, col int) {
	for row := range board {
		for col := range board[row] {
			if board[row][col] == 'K' {
				return row, col
			}
		}
	}

	return
}
