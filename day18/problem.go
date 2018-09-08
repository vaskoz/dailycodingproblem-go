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
