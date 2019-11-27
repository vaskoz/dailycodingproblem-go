package day373

import "sort"

// LongestConsecutiveSequenceBrute runs in O(N log N)
// due to sorting the input.
func LongestConsecutiveSequenceBrute(nums []int) int {
	sort.Ints(nums)

	longest := 0
	current := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			current++
		} else {
			if current > longest {
				longest = current
			}
			current = 1
		}
	}

	if current > longest {
		longest = current
	}

	return longest
}
