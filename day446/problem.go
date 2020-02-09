package day446

// IsPowerFourBrute checks if a number is a power of 4 by repeated
// division.
func IsPowerFourBrute(n int32) bool {
	if n == 0 {
		return false
	}

	for n != 1 {
		if n%4 != 0 {
			return false
		}

		n /= 4
	}

	return true
}

// IsPowerFourFaster checks if a number is a power of 4
// by looking for the highest set bits.
func IsPowerFourFaster(n int32) bool {
	return n != 0 && ((n & (n - 1)) == 0) && (n&0xAAAAAAA) == 0
}
