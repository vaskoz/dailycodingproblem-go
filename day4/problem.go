package day4

// PartitionBySign groups the negative values first, then zero values,
// then positive values.
// This operation happens in-place on the provided IntSlice.
//
// Returns the index to the first positive number.
// Runs in O(N) time and O(1) space.
func partitionBySign(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[right] > 0 {
			right--
		} else if nums[left] <= 0 {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	for left < len(nums) && nums[left] <= 0 {
		left++
	}
	return left
}

// findSmallestPositiveOnly works only on a slice of positive > 0 integers.
// Runs in O(N) since it does two sequential linear scans.
func findSmallestPositiveOnly(nums []int) int {
	for i := range nums {
		abs := nums[i]
		if abs < 0 {
			abs = -abs
		}
		if index := abs - 1; index < len(nums) {
			if nums[index] > 0 {
				nums[index] = -nums[index]
			}
		}
	}

	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}

// FindSmallestMissingPositive returns the smallest positive number missing.
// Running time is O(N) with O(1) space.
// NOTE: this modifies the input array
func FindSmallestMissingPositive(nums []int) int {
	positiveStart := partitionBySign(nums)
	return findSmallestPositiveOnly(nums[positiveStart:])
}
