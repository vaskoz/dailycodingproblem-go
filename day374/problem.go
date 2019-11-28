package day374

import (
	"errors"
	"sort"
)

var (
	errNoSuchIndex    = errors.New("no index is a fixed point")
	errInputNotSorted = errors.New("input is not sorted and must be")
)

// LowestFixedPoint returns the lowest fixed point index.
// If no such index exists, it returns an error.
// Also returns an error if input is not sorted.
// Runs in O(log N) time thanks to binary search.
// Worse case is O(N/2) because I iterate down from a match
// to find the lowest fixed point.
func LowestFixedPoint(sorted []int) (int, error) {
	if !sort.IntsAreSorted(sorted) {
		return 0, errInputNotSorted
	}

	for lo, hi := 0, len(sorted)-1; lo < hi; {
		switch mid := (lo + hi) / 2; {
		case sorted[mid] == mid:
			for sorted[mid] == mid {
				mid--
			}

			return mid + 1, nil
		case sorted[mid] > mid:
			hi = mid - 1
		default:
			lo = mid + 1
		}
	}

	return 0, errNoSuchIndex
}
