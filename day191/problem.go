package day191

import "sort"

// Interval has a start and end.
type Interval struct {
	Start, End int
}

// MinRemoveNoOverlap returns the minimum number of intervals
// to remove to remove all overlaps.
// Runs in O(N log N) time.
func MinRemoveNoOverlap(intervals []Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].End < intervals[j].End
	})
	count := 0
	current := -int(^uint(0)>>1) - 1
	for _, interval := range intervals {
		if interval.Start >= current {
			count++
			current = interval.End
		}
	}
	return len(intervals) - count
}
