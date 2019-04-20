package day241

import "sort"

// HIndex returns the H-Index of the citations from an author.
// Runs in O(N log N) using O(N) extra memory because we
// don't want to mutate the original data by sorting.
func HIndex(citations []int) int {
	copied := make([]int, len(citations))
	copy(copied, citations)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i] > copied[j]
	})
	for paper, citationCount := range copied {
		if citationCount < paper+1 {
			return paper
		}
	}
	return len(copied)
}
