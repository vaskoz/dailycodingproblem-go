package day408

// MaxProfit returns the maximum possible profit given
// exactly 'k' buy/sell pairs.
// Runs in O(N*k) time and O(N*k) space.
func MaxProfit(prices []int, k int) int {
	profits := make([][]int, k+1)
	for i := range profits {
		profits[i] = make([]int, len(prices)+1)
	}

	for i := 1; i <= k; i++ {
		diff := -int(^uint(0)>>1) - 1
		for j := 1; j < len(prices); j++ {
			diff = max(diff, profits[i-1][j-1]-prices[j-1])
			profits[i][j] = max(profits[i][j-1], prices[j]+diff)
		}
	}

	return profits[k][len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
