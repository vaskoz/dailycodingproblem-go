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
func StabIntervalPoints(X []Interval) []Point {
	if len(X) == 0 {
		return nil
	}
	sort.Slice(X, func(i, j int) bool {
		if X[i].End < X[j].End {
			return true
		} else if X[i].End == X[j].End {
			return X[i].Start < X[j].Start
		}
		return false
	})
	P := []Point{X[0].End}
	last := P[0]
	for _, interval := range X {
		if interval.Start > last {
			last = interval.End
			P = append(P, last)
		}
	}
	return P
}
