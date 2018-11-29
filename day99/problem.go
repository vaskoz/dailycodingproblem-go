package day99

import "sort"

// LongestConsecutiveSequenceBrute given an unsorted array of integers,
// returns the length of the longest consecutive elements sequence.
// Runtime O(N log N)
func LongestConsecutiveSequenceBrute(nums []int) int {
	sort.Ints(nums)
	var max, count int
	count = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			count++
		} else if count > max {
			max = count
			count = 1
		}
	}
	if count > max {
		max = count
	}
	return max
}
