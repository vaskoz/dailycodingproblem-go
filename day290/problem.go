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
				tmp := []rune{missing}
				copied = append(copied[:i-1], append(tmp, copied[i+1:]...)...)
				break
			}
		}
	}
	return copied
}
