package day448

// SegregateValues takes a slice of letters and segregates them in-place.
// O(N) time and O(1) space.
func SegregateValues(letters []rune) {
	var ri, gi int

	for i, r := range letters {
		if r == 'R' {
			letters[i], letters[gi], letters[ri] = letters[gi], letters[ri], r
			gi++
			ri++
		} else if r == 'G' {
			letters[i], letters[gi] = letters[gi], r
			gi++
		}
	}
}
