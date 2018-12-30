package day128

import "fmt"

// TowerOfHanoiMoves returns a slice of strings that describe how to optimally achieve
// the goal. O(2^N) is the complexity and result size.
func TowerOfHanoiMoves(disks int) []string {
	return towerOfHanoi(disks, 1, 3, 2)
}

func towerOfHanoi(disks, from, to, extra int) []string {
	if disks == 1 {
		return []string{fmt.Sprintf("Move %d to %d", from, to)}
	}
	moves := towerOfHanoi(disks-1, from, extra, to)
	moves = append(moves, fmt.Sprintf("Move %d to %d", from, to))
	moves = append(moves, towerOfHanoi(disks-1, extra, to, from)...)
	return moves
}
