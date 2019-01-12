package day143

import "sort"

// DNFPartitionBrute has values smaller than pivot first, followed by pivot values,
// followed by values larger than pivot last.
// Runs in O(N log N) because it just sorts.
func DNFPartitionBrute(lst []int, pivot int) {
	sort.Ints(lst)
}

// DNFPartition has values smaller than pivot first, followed by pivot values,
// followed by values larger than pivot last.
// Runs in O(N) because it just sorts.
func DNFPartition(lst []int, pivot int) {
	var i, j int
	n := len(lst) - 1
	for j <= n {
		if lst[j] < pivot {
			lst[i], lst[j] = lst[j], lst[i]
			i++
			j++
		} else if lst[j] > pivot {
			lst[j], lst[n] = lst[n], lst[j]
			n--
		} else {
			j++
		}
	}
}
