package day380

// DivisionBrute returns the quotient and remainder of division
// without using the division operator using brute-force.
func DivisionBrute(dividend, divisor int) (int, int) {
	if divisor == 0 {
		panic("divide by zero")
	}

	var dividendNeg, divisorNeg bool

	if dividend < 0 {
		dividendNeg = true
		dividend = -dividend
	}

	if divisor < 0 {
		divisorNeg = true
		divisor = -divisor
	}

	var quotient int

	for dividend >= 0 {
		dividend -= divisor
		quotient++
	}

	remainder := dividend + divisor
	quotient--

	if dividendNeg {
		quotient = -quotient
		remainder = -remainder
	}

	if divisorNeg {
		quotient = -quotient
		remainder = -remainder
	}

	return quotient, remainder
}
