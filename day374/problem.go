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
// Runs in O(log N) worst case time thanks to binary search.
func LowestFixedPoint(sorted []int) (int, error) {
	if !sort.IntsAreSorted(sorted) {
		return 0, errInputNotSorted
	}

	for lo, hi := 0, len(sorted); lo <= hi; {
		switch mid := (lo + hi) / 2; {
		case sorted[mid] == mid:
			if hi-lo == 0 {
				return mid, nil
			}

			hi = mid
		case sorted[mid] < mid:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}

	return 0, errNoSuchIndex
}
