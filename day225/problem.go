package day225

import "math/bits"

// JosephusGeneral is the general case solution.
func JosephusGeneral(n, k int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	}
	if k == 2 {
		highestBitValue := 1 << uint(bits.Len64(uint64(n))-1)
		return 2*(n-highestBitValue) + 1
	}
	return josephusGeneral(n, k) + 1
}

func josephusGeneral(n, k int) int {
	if n == 1 {
		return 0
	}
	return (josephusGeneral(n-1, k) + k) % n
}
