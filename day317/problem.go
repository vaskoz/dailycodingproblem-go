package day317

// BitwiseAndInclusiveRange returns the AND of all the numbers
// between m and n inclusive.
// This runs by performing O(abs(N-M)) AND operations.
// This is a brute force solution.
func BitwiseAndInclusiveRange(m, n uint64) uint64 {
	if m > n {
		m, n = n, m
	}
	for i := m + 1; i <= n; i++ {
		m &= i
	}
	return m
}
