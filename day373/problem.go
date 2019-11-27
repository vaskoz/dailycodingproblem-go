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

// LongestConsecutiveSequenceLinear runs in O(N).
func LongestConsecutiveSequenceLinear(nums []int) int {
	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0

	for _, num := range nums {
		if _, found := set[num-1]; !found {
			found = true
			count := 0

			for ; found; count++ {
				_, found = set[num+count+1]
			}

			if count > longest {
				longest = count
			}
		}
	}

	return longest
}
