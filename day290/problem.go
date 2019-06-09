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
		for i := 1; i < lastLen; i++ {
			if copied[i] != copied[i-1] {
				missing := 'R' + 'G' + 'B' - copied[i] - copied[i-1]
				mid := []rune{missing}
				end := append(mid, copied[i+1:]...)
				copied = append(copied[:i-1], end...)
				break
			}
		}
	}
	return copied
}
