package day25

// Match returns if there is a regex match on the given input.
func Match(input, pattern string) bool {
	var ii, pi, missed int
	for ii < len(input) && pi < len(pattern) && missed < 2 {
		if pattern[pi] == '*' {
			prev := pattern[pi-1]
			var next byte
			if pi+1 < len(pattern) {
				next = pattern[pi+1]
			}
			for ii < len(input) && input[ii] != next {
				if prev != '.' && input[ii] != prev {
					return false
				}
				ii++
			}
			if ii == len(input) && pi < len(pattern)-1 {
				return false
			}
			ii--
		} else if pattern[pi] != '.' && input[ii] != pattern[pi] {
			missed++
		}
		ii++
		pi++
	}
	return ii == len(input) && pi == len(pattern)
}
