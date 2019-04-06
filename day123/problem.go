package day123

// IsANumber answers if the given string
// can be converted to a number.
func IsANumber(str string) bool {
	var seenNeg, seenDecimal, seenE, seenDigit bool
	for _, r := range str {
		switch {
		case r == '-':
			if seenNeg || seenDigit {
				return false
			}
			seenNeg = true
		case r == '.':
			if seenDecimal || !seenDigit {
				return false
			}
			seenDecimal = true
		case r == 'e' || r == 'E':
			if seenE || !seenDigit {
				return false
			}
			seenE = true
		case r >= '0' && r <= '9':
			seenDigit = true
		default:
			return false
		}
	}
	return seenDigit
}
