package day62

// CountPathsBrute counts the number of paths from upper-left to lower-right.
// Allowed movement is only down and right.
// Runs in O(2^N) time.
func CountPathsBrute(n, m int) int {
	if n < 1 || m < 1 {
		panic("invalid size. both arguments must be 1 or greater")
	} else if n == 1 || m == 1 {
		return 1
	}
	return CountPathsBrute(m-1, n) + CountPathsBrute(m, n-1)
}

// CountPathsDP is the same solution just using dynamic programming.
// Runs in O(m*n) time and O(m*n) space.
func CountPathsDP(n, m int) int {
	if n < 1 || m < 1 {
		panic("invalid size. both arguments must be 1 or greater")
	}
	counts := make([][]int, m)
	for i := range counts {
		counts[i] = make([]int, n)
		counts[i][0] = 1
	}
	for i := 0; i < n; i++ {
		counts[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			counts[i][j] = counts[i-1][j] + counts[i][j-1]
		}
	}
	return counts[m-1][n-1]
}
