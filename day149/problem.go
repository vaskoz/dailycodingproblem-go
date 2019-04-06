package day149

// SumSublist returns the sum of the sublist from start (inclusive)
// up to end (exclusive).
// Runs in O(end-start).
func SumSublist(l []int, start, end int) int {
	var sum int
	for i := start; i < end; i++ {
		sum += l[i]
	}
	return sum
}
