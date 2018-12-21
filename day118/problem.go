package day118

import "sort"

// SquareElementsSortedBrute returns the squared result of every element of the param.
// Runs in O(N log N) due to sorting the results.
func SquareElementsSortedBrute(sorted []int) []int {
	result := make([]int, len(sorted))
	for i := range sorted {
		result[i] = sorted[i] * sorted[i]
	}
	sort.Ints(result)
	return result
}
