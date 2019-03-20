package day210

// LongestCollatzSequence returns the number with the longest
// sequence.
// First result is the number, second number is the length of the sequence.
func LongestCollatzSequence(max int) (int, int) {
	var n, longest int
	for i := max; i > 0; i-- {
		count := 0
		v := i
		for v != 1 {
			if v%2 == 0 {
				v /= 2
			} else {
				v = 3*v + 1
			}
			count++
		}
		if count > longest {
			longest = count
			n = i
		}
	}
	return n, longest
}
