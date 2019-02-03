package day165

// SmallerRightCount returns the number of smaller elements
// to the right of the position.
// Runs in O(N^2) time.
func SmallerRightCount(nums []int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if v > nums[j] {
				count++
			}
		}
		result[i] = count
	}
	return result
}
