package day143

import "sort"

// PartitionBrute has values smaller than pivot first, followed by pivot values,
// followed by values larger than pivot last.
// Runs in O(N log N) because it just sorts.
func PartitionBrute(lst []int, pivot int) {
	sort.Ints(lst)
}
