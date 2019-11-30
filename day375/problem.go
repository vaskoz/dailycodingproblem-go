package day375

import "sort"

// HIndex returns the h-index for this author.
func HIndex(citations []int) int {
	copied := make([]int, len(citations))
	copy(copied, citations)

	sort.Sort(sort.Reverse(sort.IntSlice(copied)))

	for index, citation := range copied {
		if citation < index+1 {
			return index
		}
	}

	return len(copied)
}
