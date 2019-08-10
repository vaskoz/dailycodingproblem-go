package day352

// Square represents either a white or black square
// on a crossword puzzle.
type Square int

const (
	// W is a white square
	W Square = iota
	// B is a white square
	B
)

func IsValidCrossword(puzzle [][]Square) bool {
	if isValid := checkHorizontal(puzzle); !isValid {
		return false
	}
	if isValid := checkVertical(puzzle); !isValid {
		return false
	}
	return true
}

func checkHorizontal(puzzle [][]Square) bool {
	var r, c int
	for r < len(puzzle) {
		if puzzle[r][c] == B {
			c++
		} else if puzzle[r][c] == W {
			count := 0
			for c < len(puzzle[r]) && puzzle[r][c] == W {
				count++
				c++
			}
			if count < 3 {
				return false
			}
		}
		if c == len(puzzle[r]) {
			c = 0
			r++
		}
	}
	return true
}

func checkVertical(puzzle [][]Square) bool {
	var r, c int
	for c < len(puzzle[r]) {
		if puzzle[r][c] == B {
			r++
		} else if puzzle[r][c] == W {
			count := 0
			for r < len(puzzle) && puzzle[r][c] == W {
				count++
				r++
			}
			if count < 3 {
				return false
			}
		}
		if r == len(puzzle) {
			r = 0
			c++
		}
	}
	return true
}
