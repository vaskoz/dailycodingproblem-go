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

// MaxMinPairs returns the minimum and maximum values
// from the slice of integers.
// Runs in 3N/2 which is smaller than 2*(N-2)
func MaxMinPairs(nums []int) (min, max int) {
	if len(nums) == 0 {
		return 0, 0
	} else if len(nums) == 1 {
		return nums[0], nums[0]
	}
	min, max = nums[0], nums[1]
	start := 2
	if len(nums)%2 == 1 {
		min, max = nums[0], nums[0]
		start = 1
	}
	for i := start; i < len(nums); i += 2 {
		ma, mi := nums[i], nums[i+1]
		if mi > ma {
			ma, mi = mi, ma
		}
		if ma > max {
			max = ma
		}
		if mi < min {
			min = mi
		}
	}
	return min, max
}
