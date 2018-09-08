package day18

// MaxSubarraysBrute uses no extra memory, but runs in O(k*N)
func MaxSubarraysBrute(nums []int, k int) []int {
	var result []int
	for i := 0; i < len(nums)-k+1; i++ {
		max := nums[i]
		for j := i + 1; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		result = append(result, max)
	}
	return result
}

// MaxSubarraysOnePass runs through the input only once.
func MaxSubarraysOnePass(nums []int, k int) []int {
	var result []int
	var sublist []int // O(k) extra space
	for i := 0; i < len(nums); i++ {
		if i >= k {
			result = append(result, nums[sublist[0]])
		}
		for len(sublist) > 0 && sublist[0] <= i-k {
			sublist = sublist[1:]
		}
		for len(sublist) > 0 && nums[i] >= nums[sublist[len(sublist)-1]] {
			sublist = sublist[:len(sublist)-1]
		}
		sublist = append(sublist, i)
	}
	return append(result, nums[sublist[0]])
}
