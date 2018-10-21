package day61

// NaiveIntegerExponentiation returns the result of x^y by calculating it
// naively. Runs in O(y) time due to y multiplications of x.
// y must be 0 or greater otherwise will panic.
func NaiveIntegerExponentiation(x, y int) int {
	if y < 0 {
		panic("y must be 0 or greater for integer exponentiation")
	}
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
