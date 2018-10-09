package day49

// MaxSubArray returns the maximum possible sum of a contiguous sub-array.
// Runs in O(N) time and O(1) space.
func MaxSubArray(arr []int) int {
	var maxEndingHere, maxSoFar int
	for i := range arr {
		maxEndingHere = max(arr[i], maxEndingHere+arr[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

// max returns the larger of two ints.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
