package day25

// Match returns if there is a regex match on the given input.
func Match(input, pattern string) bool {
	var ii, pi, missed int
	inputR := []rune(input)
	patternR := []rune(pattern)
	for isPossible(ii, len(inputR), pi, len(patternR), missed) {
		if patternR[pi] == '*' {
			prev := patternR[pi-1]
			var next rune
			if pi+1 < len(patternR) {
				next = patternR[pi+1]
			}
			for shouldKeepLooking(ii, len(inputR), inputR, next) {
				if prev != '.' && inputR[ii] != prev {
					return false
				}
				ii++
			}
			if remainingPattern(ii, len(inputR), pi, len(patternR)) {
				return false
			}
			ii--
		} else if missedMatch(patternR[pi], inputR[ii]) {
			missed++
		}
		ii++
		pi++
	}
	return finished(ii, len(input), pi, len(pattern))
}

// finished is true if we've completed processing the input and pattern strings.
func finished(inputIndex, inputSize, patternIndex, patternSize int) bool {
	return inputIndex == inputSize && patternIndex == patternSize
}

// isPossible returns true if its possible the pattern could still match.
func isPossible(inputIndex, inputSize, patternIndex, patternSize, numMissed int) bool {
	return inputIndex < inputSize && patternIndex < patternSize && numMissed < 2
}

// missedMatch returns true if a match was missed.
func missedMatch(patternRune, inputRune rune) bool {
	return patternRune != '.' && inputRune != patternRune
}

// remainingPattern returns true if we finished processing the input, but still have pattern remaining.
func remainingPattern(inputIndex, inputSize, patternIndex, patternSize int) bool {
	return inputIndex == inputSize && patternIndex < patternSize-1
}

// shouldKeepLooking returns true if the wildcard should continue looking at the input.
func shouldKeepLooking(inputIndex, inputSize int, inputR []rune, next rune) bool {
	return inputIndex < inputSize && inputR[inputIndex] != next
}
