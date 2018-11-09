package day79

// CouldBeNonDecreasingByOneChange answers the question:
// whether the array could become non-decreasing by modifying
// at most 1 element.
// Runs in O(N) time.
func CouldBeNonDecreasingByOneChange(nums []int) bool {
	p := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if p != -1 {
				return false
			}
			p = i - 1
		}
	}
	return p == -1 ||
		p == 0 ||
		p == len(nums)-2 ||
		nums[p-1] <= nums[p+1] ||
		nums[p] <= nums[p+2]
}
