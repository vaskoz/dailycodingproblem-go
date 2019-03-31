package day219

import "fmt"

// Possible winning positions.
// Vertical: 7 columns * 3 positions per column = 21 vertical winning positions.
// Horizontal: 6 rows * 4 positions per column = 24 horizontal winning positions.
// DiagUp: 12
// DiagDown: 12
// Total == 69 winning positions checking 4 positions == 276 checks

// Player represents a ConnectFour player or none.
type Player int

const (
	// None is neither player.
	None Player = iota
	// Red is a position owned by the 'red' player.
	Red
	// Black is a position owned by the 'black' player.
	Black
)

// Column is zero-based from left to right.
type Column int

// Row is zero-based from bottom to top.
type Row int

// ConnectFourBoard represents a column-major board since
// moves are played by column.
// 7 vertical columns and 6 horizontal rows.
type ConnectFourBoard struct {
	board map[Column]map[Row]Player
}

// Move is a piece played by a player.
// Player must be either 'Red' or 'Black' to play a valid move.
func (cfb *ConnectFourBoard) Move(player Player, col Column) error {
	if player != Red && player != Black {
		return fmt.Errorf("must play either Red or Black, you tried %v", player)
	}
	if cfb.board == nil {
		cfb.board = make(map[Column]map[Row]Player)
	}
	valid := false
	for row := Row(0); row < 6; row++ {
		if cfb.board[col] == nil {
			cfb.board[col] = make(map[Row]Player)
		}
		if _, found := cfb.board[col][row]; !found {
			cfb.board[col][row] = player
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("column full. can't play in this column")
	}
	return nil
}

// Winner returns the winner of the current board.
// If there is no winner yet, then Player=None is returned.
func (cfb *ConnectFourBoard) Winner() Player {
	if p := cfb.checkVerticalWinner(); p != None {
		return p
	}
	if p := cfb.checkHorizontalWinner(); p != None {
		return p
	}
	return None
}

func (cfb *ConnectFourBoard) checkHorizontalWinner() Player {
	for row := Row(0); row < 6; row++ {
		for col := Column(0); col < 4; col++ {
			if p, found := cfb.board[col][row]; found {
				win := true
				for c := col + 1; c < col+4; c++ {
					if prow, found := cfb.board[c][row]; !found || prow == p {
						win = false
						break
					}
				}
				if win {
					return p
				}
			}
		}
	}
	return None
}

func (cfb *ConnectFourBoard) checkVerticalWinner() Player {
	for col := Column(0); col < 7; col++ {
		for row := Row(5); row > 2; row-- {
			if p, found := cfb.board[col][row]; found {
				win := true
				for r := row - 1; r > row-4; r-- {
					if pcol, found := cfb.board[col][r]; !found || pcol == p {
						win = false
						break
					}
				}
				if win {
					return p
				}
			}
		}
	}
	return None
}
