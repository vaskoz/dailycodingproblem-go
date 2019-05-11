package day257

// SmallestWindowThatMustBeSorted returns the smallest window that must be
// sorted to result in a completely sorted slice.
func SmallestWindowThatMustBeSorted(nums []int) (start, end int) {
	maxSoFar := -int(^uint(0)>>1) - 1
	for i := range nums {
		if maxSoFar < nums[i] {
			maxSoFar = nums[i]
		}
		if nums[i] < maxSoFar {
			end = i
		}
	}
	minSoFar := int(^uint(0) >> 1)
	for i := range nums {
		j := len(nums) - 1 - i
		if minSoFar > nums[j] {
			minSoFar = nums[j]
		}
		if nums[j] > minSoFar {
			start = j
		}
	}
	return
}
