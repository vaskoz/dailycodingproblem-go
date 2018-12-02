package day102

// ContiguousSumBrute uses brute-force to return the
// contiguous subset that sums to K.
// Runtime is O(N^2) and O(1) space.
func ContiguousSumBrute(nums []int, k int) []int {
	var result []int
	for i := range nums {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				return nums[i : j+1]
			}
		}
	}
	return result
}

// ContiguousSumFaster uses a window to return the
// contiguous subset that sums to K.
// Runtime is O(N) and O(1) space.
func ContiguousSumFaster(nums []int, k int) []int {
	var result []int
	var sum, begin int
	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		if sum == k {
			result = nums[begin : end+1]
			break
		} else if sum > k {
			sum -= nums[begin]
			sum -= nums[end]
			begin++
			end--
		}
	}
	return result
}
