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
	height := Row(6)
	width := Column(7)
	for row := Row(0); row < height; row++ {
		for col := Column(0); col < width; col++ {
			player := cfb.board[col][row]
			if player == None {
				continue
			}
			if col+3 < width &&
				player == cfb.board[col+1][row] &&
				player == cfb.board[col+2][row] &&
				player == cfb.board[col+3][row] {
				return player
			}
			if row+3 < height {
				if player == cfb.board[col][row+1] &&
					player == cfb.board[col][row+2] &&
					player == cfb.board[col][row+3] {
					return player
				}
				if col+3 < width &&
					player == cfb.board[col+1][row+1] &&
					player == cfb.board[col+2][row+2] &&
					player == cfb.board[col+3][row+3] {
					return player
				}
				if col-3 >= 0 &&
					player == cfb.board[col-1][row+1] &&
					player == cfb.board[col-2][row+2] &&
					player == cfb.board[col-3][row+3] {
					return player
				}
			}
		}
	}
	return None
}
