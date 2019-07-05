package day317

import (
	"math/bits"
)

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

// BitwiseAndInclusiveRangeFaster returns the AND of all the numbers
// between m and n inclusive.
// This is a smarter bit manipulation version.
func BitwiseAndInclusiveRangeFaster(m, n uint64) uint64 {
	var result uint64
	for {
		mMsb := bits.Len64(m)
		nMsb := bits.Len64(n)
		if mMsb != nMsb || mMsb == 0 {
			break
		}
		val := uint64(1 << (uint(mMsb) - 1))
		result += val
		m -= val
		n -= val
	}
	return result
}
