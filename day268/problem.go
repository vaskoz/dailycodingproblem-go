package day268

// IsPowerFourBrute checks if a number is a power of 4 by repeated
// division.
func IsPowerFourBrute(n int) bool {
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
