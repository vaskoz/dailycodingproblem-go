package day99

import "sort"

// LongestConsecutiveSequenceBrute given an unsorted array of integers,
// returns the length of the longest consecutive elements sequence length.
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

// LongestConsecutiveSequenceLinear given an unsorted array of integers,
// returns the length of the longest consecutive elements sequence length.
// Runtime O(N)
func LongestConsecutiveSequenceLinear(nums []int) int {
	set := make(map[int]struct{})
	var max int
	for _, val := range nums {
		set[val] = struct{}{}
	}
	for _, val := range nums {
		if _, found := set[val-1]; !found {
			found = true
			var count int
			for count = 1; found; count++ {
				_, found = set[val+count]
			}
			if count-1 > max {
				max = count - 1
			}
		}
	}
	return max
}
