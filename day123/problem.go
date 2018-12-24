package day123

// IsANumber answers if the given string
// can be converted to a number.
func IsANumber(str string) bool {
	var seenNeg, seenDecimal, seenE, seenDigit bool
	for _, r := range str {
		if r == '-' {
			if seenNeg || seenDigit {
				return false
			}
			seenNeg = true
		} else if r == '.' {
			if seenDecimal || !seenDigit {
				return false
			}
			seenDecimal = true
		} else if r == 'e' || r == 'E' {
			if seenE || !seenDigit {
				return false
			}
			seenE = true
		} else if r >= '0' && r <= '9' {
			seenDigit = true
		} else {
			return false
		}
	}
	return seenDigit
}
