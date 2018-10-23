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
