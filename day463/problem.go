package day463

// MinimizeQux returns the smallest Qux slice
// possible.
// Runs in O(N*K) where K is the number of elements that can removed.
// Uses O(N) extra space to avoid modifying the input.
func MinimizeQux(input []rune) []rune {
	lastLen := 0
	copied := make([]rune, len(input))
	copy(copied, input)

	for len(copied) != lastLen {
		lastLen = len(copied)
		copied = minimizeQux(copied)
	}

	return copied
}

func minimizeQux(input []rune) []rune {
	min := input

	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			missing := 'R' + 'G' + 'B' - input[i] - input[i-1]
			end := append([]rune{missing}, input[i+1:]...)
			candidate := append([]rune{}, input[:i-1]...)
			candidate = append(candidate, end...)
			candidate = minimizeQux(candidate)

			if len(candidate) < len(min) {
				min = candidate
			}
		}
	}

	return min
}
