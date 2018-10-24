package day63

import "reflect"

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
			if reflect.DeepEqual(target, puzzle[row][col:col+len(target)]) {
				return true
			}
			if coll-len(target)-row >= 0 {
				vertical := true
				for i := range target {
					if puzzle[i+row][col] != target[i] {
						vertical = false
					}
				}
				if vertical {
					return true
				}
			}
		}
	}
	return false
}
