package day248

// MaxIf is a typical max function using an if branch.
func MaxIf(a, b int) int {
	if a > b {
		return a
	}
	return b
}
