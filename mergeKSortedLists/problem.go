package mergeklist

import "math"

// MergeKSortedLists merges K-sorted lists into a single list
// Runtime is O(K*(K*N)) = O(K^2*N) assuming each list of is length N
// Space is O(K) for the indicies without considering the answer list of O(K*N) size.
// Use this if K is very small
func MergeKSortedLists(lists [][]int) []int {
	idxs := make([]int64, len(lists))
	var result []int
	hasMore := true
	for hasMore {
		listID := -1
		smallest := math.MaxInt64
		// find the smallest next value
		for i, idx := range idxs {
			if idx < int64(len(lists[i])) && lists[i][idx] < smallest {
				smallest = lists[i][idx]
				listID = i
			}
		}
		if listID == -1 {
			hasMore = false
		} else {
			result = append(result, smallest)
			idxs[listID]++
		}
	}
	return result
}
