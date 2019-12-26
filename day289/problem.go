package day289

// CanFirstPlayerForceWinNimGame answers if the first player
// can force a win with the given heap configuration.
func CanFirstPlayerForceWinNimGame(heaps []int) bool {
	sum := 0

	for _, heap := range heaps {
		sum ^= heap
	}

	return sum != 0
}
