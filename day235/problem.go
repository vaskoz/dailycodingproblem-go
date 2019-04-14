package day235

// MaxMinNotGoodEnough returns the minimum and maximum values
// from the slice of integers.
// Runs in 2*(N-1) because it does N-1 comparisons for both
// min and max.
func MaxMinNotGoodEnough(nums []int) (min, max int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return min, max
}
