package day350

import "math"

// SmallestNumberOfPerfectSquaresSum returns the number
// of perfect squares that sum up to N.
// Runs in O(2^N) time.
func SmallestNumberOfPerfectSquaresSum(n int) int {
	if n < 4 {
		return n
	}
	res := n
	for i := 1; i <= n; i++ {
		tmp := i * i
		if tmp > n {
			break
		} else {
			res = min(res, 1+SmallestNumberOfPerfectSquaresSum(n-tmp))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SmallestNumberOfPerfectSquaresSumFaster returns the number
// of perfect squares that sum up to N.
// Runs in O(N) time and uses O(N) space for dynamic programming.
func SmallestNumberOfPerfectSquaresSumFaster(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < 4; i++ {
		dp[i] = i
	}
	for i := 4; i <= n; i++ {
		dp[i] = i
		for x := 1; x <= int(math.Ceil(math.Sqrt(float64(i)))); x++ {
			tmp := x * x
			if tmp > i {
				break
			} else {
				dp[i] = min(dp[i], 1+dp[i-tmp])
			}
		}
	}
	return dp[n]
}
