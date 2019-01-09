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

// SingleOccurrenceLinear returns numbers that occur only once.
// Runs in O(N) time and O(1) space.
// Only works if the repeating elements only repeat twice and
// if only 2 numbers occur once.
func SingleOccurrenceLinear(nums []int) []int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return nums
	}
	xor := nums[0]
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}
	results := make([]int, 2)
	rightMost := xor & ^(xor - 1)
	for i := range nums {
		if nums[i]&rightMost != 0 {
			results[0] ^= nums[i]
		} else {
			results[1] ^= nums[i]
		}
	}
	return results
}
