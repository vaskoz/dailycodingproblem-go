package day338

// NextIntSameSetBits returns the next greater
// integer with the same set bits.
func NextIntSameSetBits(n int) int {
	res := 0
	if n != 0 {
		rightMostBitInt := n & -n
		nextHigherSetBitInt := n + rightMostBitInt
		d := n ^ nextHigherSetBitInt
		d /= rightMostBitInt
		d >>= 2
		res = nextHigherSetBitInt | d
	}
	return res
}
