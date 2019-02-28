package day190

// MaxWrapSubarraySum returns the maximum possible sum for
// a subarray in a circular array.
// Runs in O(N) time and O(N) space because the data is mutated.
func MaxWrapSubarraySum(nums []int) int {
	data := make([]int, len(nums))
	copy(data, nums)
	maxKadane := KadaneAlg(data)
	maxWrap := 0
	for i, v := range data {
		maxWrap += v
		data[i] = -data[i]
	}
	maxWrap += KadaneAlg(data)
	return max(maxWrap, maxKadane)
}

// KadaneAlg implements https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm
func KadaneAlg(nums []int) int {
	maxEndingHere, maxSoFar := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
