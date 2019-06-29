package day311

// PeakIndexBrute returns the index of a peak.
// Runs in O(N).
func PeakIndexBrute(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := range nums {
		switch i {
		case 0:
			if nums[i] > nums[i+1] {
				return i
			}
		case len(nums) - 1:
			if nums[i] > nums[i-1] {
				return i
			}
		default:
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return -1
}
