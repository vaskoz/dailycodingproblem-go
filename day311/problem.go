package day311

// PeakIndex returns the index of a peak.
// Runs in O(log N).
func PeakIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return 0
	}
	return peakIndex(nums, 0, len(nums)-1)
}

func peakIndex(nums []int, low, high int) int {
	mid := (low + high) / 2
	if (mid == 0 || nums[mid-1] <= nums[mid]) &&
		(mid == len(nums)-1 || nums[mid+1] <= nums[mid]) {
		return mid
	} else if mid > 0 && nums[mid-1] > nums[mid] {
		return peakIndex(nums, low, mid-1)
	}
	return peakIndex(nums, mid+1, high)
}

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
