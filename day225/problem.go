package day225

import "math/bits"

// JosephusGeneral is the general case solution.
func JosephusGeneral(n, k int) int {
	if k == 2 {
		return josephusGeneralK2(n)
	}
	return josephusGeneral(n, k) + 1
}

func josephusGeneralK2(n int) int {
	highestBit := bits.Len64(uint64(n))
	highestBit = 1 << uint(highestBit-1)
	return 2*(n-highestBit) + 1
}

func josephusGeneral(n, k int) int {
	if n == 1 {
		return 0
	}
	return (josephusGeneral(n-1, k) + k) % n
}
