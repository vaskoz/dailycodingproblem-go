package day269

// SimulateDominoes runs a simulation of dominoes falling
// based on an initial state. It returns the end state.
func SimulateDominoes(initial string) string {
	start := make([]rune, len([]rune(initial)))
	copy(start, []rune(initial))
	result := make([]rune, len(start))
	changed := true
	for changed {
		changed = false
		for i, domino := range start {
			if domino == '.' {
				switch {
				case i != 0 && i != len(start)-1 && start[i-1] == 'R' && start[i+1] == 'L':
					result[i] = start[i]
				case i != 0 && start[i-1] == 'R':
					result[i] = 'R'
					changed = true
				case i != len(start)-1 && start[i+1] == 'L':
					result[i] = 'L'
					changed = true
				default:
					result[i] = start[i]
				}
			} else {
				result[i] = start[i]
			}
		}
		copy(start, result)
	}
	return string(result)
}
