package day47

// MaxProfitBrute calculates the maximum profit from these chronological stock prices.
// Runs in O(N^2) and O(1) space.
func MaxProfitBrute(stockPrices []int) int {
	var maxProfit int
	for i := range stockPrices {
		for j := i + 1; j < len(stockPrices); j++ {
			if profit := stockPrices[j] - stockPrices[i]; profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
