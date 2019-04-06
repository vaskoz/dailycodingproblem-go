package day200

import "sort"

// Point is an integer representing a point on a line.
type Point int

// Interval on a line from start to end inclusive at both ends.
type Interval struct {
	Start, End Point
}

// StabIntervalPoints computes the smallest set of points that stabs X.
// Runs in O(N log N), but modifies the input slice by sorting.
func StabIntervalPoints(x []Interval) []Point {
	if len(x) == 0 {
		return nil
	}
	sort.Slice(x, func(i, j int) bool {
		if x[i].End < x[j].End {
			return true
		} else if x[i].End == x[j].End {
			return x[i].Start < x[j].Start
		}
		return false
	})
	p := []Point{x[0].End}
	last := p[0]
	for _, interval := range x {
		if interval.Start > last {
			last = interval.End
			p = append(p, last)
		}
	}
	return p
}
