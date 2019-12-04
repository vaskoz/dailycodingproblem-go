package day380

// DivisionBrute returns the quotient and remainder of division
// without using the division operator using brute-force.
func DivisionBrute(dividend, divisor int) (int, int) {
	var quotient int

	for dividend >= 0 {
		dividend -= divisor
		quotient++
	}

	remainder := dividend + divisor
	quotient--

	return quotient, remainder
}
