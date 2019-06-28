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
