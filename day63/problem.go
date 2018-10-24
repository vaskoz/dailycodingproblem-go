package day63

// FindWordLD answers if the target word is found in the puzzle.
// Only searches left-to-right and top-to-bottom directions.
// Brute force runtime where N=rows, M=cols, and K=len(target)
// Runtime is O((N-K)*(M-K)*K)
func FindWordLD(puzzle [][]rune, target []rune) bool {
	coll := len(puzzle)
	for row := range puzzle {
		rowl := len(puzzle[row])
		for col := range puzzle[row] {
			if rowl-len(target)-col < 0 {
				break
			}
			if checkLeftToRight(puzzle, row, col, target) {
				return true
			}
			if coll-len(target)-row >= 0 {
				if checkTopToDown(puzzle, row, col, target) {
					return true
				}
			}
		}
	}
	return false
}

func checkLeftToRight(puzzle [][]rune, row, col int, target []rune) bool {
	horizontal := true
	for i, r := range target {
		if r != puzzle[row][i+col] {
			horizontal = false
			break
		}
	}
	return horizontal
}

func checkTopToDown(puzzle [][]rune, row, col int, target []rune) bool {
	vertical := true
	for i, r := range target {
		if puzzle[i+row][col] != r {
			vertical = false
			break
		}
	}
	return vertical
}
