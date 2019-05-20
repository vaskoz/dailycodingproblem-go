package day271

import "sort"

// BinarySearch is a handwritten binary search.
func BinarySearch(sorted []int, x int) int {
	left, right := 0, len(sorted)-1
	for left < right {
		mid := (left + right) / 2
		val := sorted[mid]
		switch {
		case x == val:
			return mid
		case x < val:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return -1
}

// BinarySearchSortPkg delegates to the sort package.
func BinarySearchSortPkg(sorted []int, x int) int {
	index := sort.SearchInts(sorted, x)
	if sorted[index] == x {
		return index
	}
	return -1
}
