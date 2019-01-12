package day143

import "sort"

// PartitionBrute has values smaller than pivot first, followed by pivot values,
// followed by values larger than pivot last.
// Runs in O(N log N) because it just sorts.
func PartitionBrute(lst []int, pivot int) {
	sort.Ints(lst)
}

// Partition has values smaller than pivot first, followed by pivot values,
// followed by values larger than pivot last.
// Runs in O(N) because it just sorts.
func Partition(lst []int, pivot int) {
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
