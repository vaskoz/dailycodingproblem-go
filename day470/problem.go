package day470

// ClosestLargerNumberIndexBrute looks for the closest number larger
// than the target and returns the index of that number.
// If no number if found, then -1 is returned.
// Runs in O(N) time.
func ClosestLargerNumberIndexBrute(nums []int, targetIndex int) int {
	if targetIndex < 0 || targetIndex >= len(nums) {
		panic("index out of bounds")
	}

	val := nums[targetIndex]
	left, right := targetIndex-1, targetIndex+1

	for left >= 0 || right < len(nums) {
		if left >= 0 && nums[left] > val {
			return left
		} else if right < len(nums) && nums[right] > val {
			return right
		}
		left--
		right++
	}

	return -1
}

// PreprocessClosestLargerNumberIndex pre-processes the input and returns a new
// slice that can answer the question for any index in O(1).
// This pre-processing takes O(N^2) time.
func PreprocessClosestLargerNumberIndex(nums []int) []int {
	res := make([]int, len(nums))

	for i := range nums {
		res[i] = ClosestLargerNumberIndexBrute(nums, i)
	}

	return res
}
