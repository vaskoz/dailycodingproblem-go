package day68

// BoardPosition represents a row and column on a square chessboard.
type BoardPosition struct {
	Row, Col int
}

// CountAttackingBishopPairs returns the count of the number of attacking
// pairs of Bishops.
// M is the dimension of the chessboard. M x M
// Runtime is O(N^2) where N is the number of bishops.
func CountAttackingBishopPairs(bishops []BoardPosition) int {
	count := 0
	for i := 0; i < len(bishops)-1; i++ {
		first := bishops[i]
		for j := i + 1; j < len(bishops); j++ {
			second := bishops[j]
			rowdiff := second.Row - first.Row
			coldiff := second.Col - first.Col
			if rowdiff == coldiff || rowdiff == -coldiff {
				count++
			}
		}
	}
	return count
}
