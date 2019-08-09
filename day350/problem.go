package day350

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
