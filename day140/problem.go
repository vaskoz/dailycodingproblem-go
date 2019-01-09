package day140

import (
	"sort"
)

// SingleOccurrenceBrute returns numbers that occur only once.
// Runs in O(N log N) time and O(1) space.
func SingleOccurrenceBrute(nums []int) []int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return nums
	}
	sort.Ints(nums)
	var results []int
	last := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != last && count == 1 {
			results = append(results, nums[i-1])
			last = nums[i]
			count = 0
		} else if nums[i] != last {
			last = nums[i]
			count = 0
		}
		count++
	}
	return results
}
