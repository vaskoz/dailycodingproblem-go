package day77

import "sort"

// Interval represents a linear span between start and end positions.
type Interval struct {
	Start, End int
}

// MergeOverlappingIntervals merges any overlapping intervals.
// Runs in O(N log N) because of the sorting.
func MergeOverlappingIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		switch {
		case intervals[i].Start < intervals[j].Start:
			return true
		case intervals[i].Start == intervals[j].Start:
			return intervals[i].End <= intervals[j].End
		default:
			return false
		}
	})
	merged := make([]Interval, 0, len(intervals))
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < current.End {
			if intervals[i].End > current.End {
				current.End = intervals[i].End
			}
		} else {
			merged = append(merged, current)
			current = intervals[i]
		}
	}
	merged = append(merged, current)
	return merged
}
