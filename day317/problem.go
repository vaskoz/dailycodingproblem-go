package day317

// BitwiseAndInclusiveRange returns the AND of all the numbers
// between m and n inclusive.
// This runs by performing O(abs(N-M)) AND operations.
// This is a brute force solution.
func BitwiseAndInclusiveRange(m, n uint64) uint64 {
	var small, large uint64
	if m < n {
		small = m
		large = n
	} else {
		small = n
		large = m
	}
	result := small
	for i := small + 1; i <= large; i++ {
		result &= i
	}
	return result
}
