package day144

// ClosestLargerNumberIndexBrute looks for the closest number larger
// than the target and returns the index of that number.
// If no number if found, then -1 is returned.
// Runs in O(N) time.
func ClosestLargerNumberIndexBrute(nums []int, targetIndex int) int {
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
