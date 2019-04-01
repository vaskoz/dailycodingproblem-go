package day220

// AlternatingCoinsFirst solves the game using dynamic programming.
// You and an opponent take turns choosing either the first or
// last coin from the row, removing it from the row, and receiving the value of the coin.
func AlternatingCoinsFirst(coins []int) int {
	value := make([][]int, len(coins))
	for i := range value {
		value[i] = make([]int, len(coins))
	}
	for interval := range value {
		for i, j := 0, interval; j < len(coins); i, j = i+1, j+1 {
			var a, b, c int
			if i+2 <= j {
				a = value[i+2][j]
			}
			if i+1 <= j-1 {
				b = value[i+1][j-1]
			}
			if i <= j-2 {
				c = value[i][j-2]
			}
			value[i][j] = max(coins[i]+min(a, b), coins[j]+min(b, c))
		}
	}
	return value[0][len(coins)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
