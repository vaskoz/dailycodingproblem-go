package day224

// SmallestPositiveIntNotInSubsetSum returns the
// smallest positive integer that is not the sum of a subset of the array,
// given a sorted array.
// Runs in O(N) time and O(1) space.
func SmallestPositiveIntNotInSubsetSum(nums []int) int {
	result := 1
	for i := 0; i < len(nums) && nums[i] <= result; i++ {
		result += nums[i]
	}
	return result
}
