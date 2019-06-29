package day310

import "math/bits"

// CountSetMathBits returns the number of set bits (1).
// It uses math/bits.
func CountSetMathBits(num uint64) []int {
	result := make([]int, 0, int(num))
	for i := uint64(1); i <= num; i++ {
		result = append(result, bits.OnesCount64(i))
	}
	return result
}

// PopcntA is a manual implementation.
func PopcntA(num uint64) []int {
	result := make([]int, 0, int(num))
	for i := uint64(1); i <= num; i++ {
		x := i
		x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
		x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
		x = (x & 0x0f0f0f0f0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f0f0f0f0f)
		x = (x & 0x00ff00ff00ff00ff) + ((x >> 8) & 0x00ff00ff00ff00ff)
		x = (x & 0x0000ffff0000ffff) + ((x >> 16) & 0x0000ffff0000ffff)
		x = (x & 0x00000000ffffffff) + ((x >> 32) & 0x00000000ffffffff)
		result = append(result, int(x))
	}
	return result
}

// PopcntB is a second manual implementation.
func PopcntB(num uint64) []int {
	result := make([]int, 0, int(num))
	for i := uint64(1); i <= num; i++ {
		x := i
		x -= (x >> 1) & 0x5555555555555555
		x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
		x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
		x += x >> 8
		x += x >> 16
		x += x >> 32
		result = append(result, int(x))
	}
	return result
}
