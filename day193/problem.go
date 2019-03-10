package day193

// MaxProfitWithFee returns the maximum possible profit
// by making transactions where each one has a fee.
// A backtracking algorithm that tries all profitable paths
// looking for the most profitable.
func MaxProfitWithFee(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > buy+fee {
			thisTradeProfit := prices[i] - buy - fee
			totalProfit := thisTradeProfit + MaxProfitWithFee(prices[i+1:], fee)
			if totalProfit > maxProfit {
				maxProfit = totalProfit
			}
		}
	}
	return maxProfit
}
