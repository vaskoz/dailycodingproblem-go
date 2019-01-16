package day147

import (
	"sort"
	_ "unsafe" // nolint
)

// Reverse conforms to the interface asked in the problem.
func Reverse(data sort.Interface, i, j int) {
	SortRange(sort.Reverse(data), i, j)
}

// SortRange extends the abilities of sort.Sort by allowing
// a subset of the full range to be sorted.
// For this example, I'm gonna reach into the stdlib "sort" package,
// but this is fragile because it reuses unexported functions.
func SortRange(data sort.Interface, i, j int) {
	n := data.Len()
	if i < 0 || i >= n || j < 0 || j >= n {
		panic("Indicies out of range")
	}
	if i >= j {
		panic("i must be before j")
	}
	quickSort(data, i, j, maxDepth(j-i))
}

//go:linkname quickSort sort.quickSort
func quickSort(data sort.Interface, a, b, maxDepth int)

//go:linkname maxDepth sort.maxDepth
func maxDepth(n int) int
