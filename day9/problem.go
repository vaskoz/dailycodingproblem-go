package day9

// MaximumNonAdjacentSum calculates the largest possible sum without using adjacent values.
// Runs in O(N) time and O(1) space.
func MaximumNonAdjacentSum(nums []int) int {
	var one, two, result int
	for i, val := range nums {
		switch {
		case i == 0:
			result = val
		case i == 1:
			result = max(result, val)
		default:
			result = max(one, val+two)
		}
		two = one
		one = result
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
