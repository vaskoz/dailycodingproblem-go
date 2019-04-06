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

// MaxProfit calculates the maximum profit from these chronological stock prices.
// Runs in O(N) time and O(1) space.
func MaxProfit(stockPrices []int) int {
	if len(stockPrices) == 0 {
		return 0
	}
	var maxProfit, lowestPrice int
	lowestPrice = stockPrices[0]
	for _, price := range stockPrices {
		if price < lowestPrice {
			lowestPrice = price
		} else if profit := price - lowestPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
