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
	var count int

	for i, b1 := range bishops {
		for _, b2 := range bishops[i+1:] {
			if r, c := b2.Row-b1.Row, b2.Col-b1.Col; r == c || r == -c {
				count++
			}
		}
	}

	return count
}
