package day464

import (
	"sort"
)

// LargestSubsetPairs finds the largest subset
// such that every pair of elements in the subset (i, j)
// satisfies either i % j = 0 or j % i = 0.
// Runs in O(N^2) time
func LargestSubsetPairs(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	dp := make([]int, len(nums))
	dp[len(dp)-1] = 1
	largestI, largestVal := 0, 0

	for i := len(dp) - 2; i >= 0; i-- {
		m := 0

		for j := i + 1; j < len(dp); j++ {
			if nums[j]%nums[i] == 0 {
				m = max(m, dp[j])
			}
		}

		if dp[i] = 1 + m; dp[i] > largestVal {
			largestI = i
			largestVal = dp[i]
		}
	}

	result := make([]int, 0, largestVal)

	for i := largestI; i < len(nums); i++ {
		if nums[i]%nums[largestI] == 0 {
			result = append(result, nums[i])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
