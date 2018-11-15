package day86

// MinRemovedParens calculates the minimum of parentheses to remove to make
// a list of them valid. Each open must have a corresponding close.
// Runs in O(N) time with O(1) space.
func MinRemovedParens(parens string) int {
	opens := 0
	count := 0
	for _, r := range parens {
		if r == '(' {
			opens++
		} else if r == ')' {
			if opens == 0 {
				count++
			} else {
				opens--
			}
		}
	}
	return count + opens
}
