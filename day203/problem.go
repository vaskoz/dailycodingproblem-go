package day203

// FindSmallestOnRotated is binary search for the smallest value
// on a sorted but rotated slice.
// Runs in O(log N) time and O(1) space.
func FindSmallestOnRotated(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
