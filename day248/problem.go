package day248

// MaxIf is a typical max function using an if branch.
func MaxIf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxSubtractAndShift finds the max using subtraction
// and right shifting.
func MaxSubtractAndShift(a, b int) int {
	return a - ((a - b) & ((a - b) >> 63))
}
