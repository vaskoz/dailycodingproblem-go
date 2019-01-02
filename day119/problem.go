package day119

import "sort"

// Interval has a start and an end.
type Interval struct {
	Start, End int
}

// SmallestSetCoverage returns the smallest intervals that cover
// the given intervals.
func SmallestSetCoverage(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start ||
			(intervals[i].Start == intervals[j].Start &&
				intervals[i].End < intervals[j].End)
	})
	var results []Interval
	highMin := 0
	lowMax := intervals[0].End
	highest := lowMax
	count := 1
	n := len(intervals)
	for i := 1; i <= n; i++ {
		if i < n && intervals[i].Start <= highest {
			highest = max(intervals[i].End, highest)
			highMin = max(intervals[i].Start, highMin)
			lowMax = min(lowMax, intervals[i].End)
			highMin = max(highMin, lowMax)
			count++
		} else {
			if count == 1 {
				one := intervals[i-1].Start
				lowMax = one
				highMin = one
			}
			results = append(results, Interval{lowMax, highMin})
			if i < n {
				lowMax = intervals[i].End
				highest = intervals[i].End
				count = 1
			}
		}
	}
	return results
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
