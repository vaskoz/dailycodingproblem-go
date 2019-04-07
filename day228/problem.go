package day228

import (
	"sort"
	"strconv"
	"strings"
)

// LargestPossibleInteger returns the largest possible integer
// constructed by these given numbers.
func LargestPossibleInteger(nums []int) int {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		first, _ := strconv.Atoi(strs[i] + strs[j])
		second, _ := strconv.Atoi(strs[j] + strs[i])
		return first > second
	})
	result, _ := strconv.Atoi(strings.Join(strs, ""))
	return result
}
