package day290

// MinimizeQux returns the smallest Qux slice
// possible.
// Runs in O(N*K) where K is the number of elements that can removed.
// Uses O(N) extra space to avoid modifying the input.
func MinimizeQux(input []rune) []rune {
	copied := make([]rune, len(input))
	copy(copied, input)
	var lastLen int
	for len(copied) != lastLen {
		lastLen = len(copied)
		copied = minimizeQux(copied)
	}
	return copied
}

func minimizeQux(input []rune) []rune {
	copied := make([]rune, len(input))
	copy(copied, input)
	min := copied
	for i := 1; i < len(copied); i++ {
		if copied[i] != copied[i-1] {
			missing := 'R' + 'G' + 'B' - copied[i] - copied[i-1]
			end := append([]rune{missing}, copied[i+1:]...)
			candidate := append([]rune{}, copied[:i-1]...)
			candidate = append(candidate, end...)
			candidate = minimizeQux(candidate)
			if len(candidate) < len(min) {
				min = candidate
			}
		}
	}
	return min
}
