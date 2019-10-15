package day357

// StringTreeDepth returns the height of the binary tree
// using the problem stated notation for the tree.
// Runs in O(N) time where N is the length of the input string
//  and O(1) space.
func StringTreeDepth(tree string) int {
	var maxStreak, current int

	for _, r := range tree {
		switch r {
		case '(':
			current++
		case ')':
			current--
		}

		if current > maxStreak {
			maxStreak = current
		}
	}

	return maxStreak
}
