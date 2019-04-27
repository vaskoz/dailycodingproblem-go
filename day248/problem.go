package day248

// MaxIf is a typical max function using an if branch.
func MaxIf(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MaxSubtractAndShift finds the max using subtraction
// and right shifting.
func MaxSubtractAndShift(a, b int64) int64 {
	return a - ((a - b) & ((a - b) >> 63))
}

// MaxXor uses xor and branching to calculate the max.
// Go requires if branching to convert a bool to an int64.
// NOTE: No benefit here due to branching if, might as well use MaxIf.
// Benchmark perf shows actually worse than MaxIf as expected.
func MaxXor(a, b int64) int64 {
	bit := int64(0)
	if a < b {
		bit = int64(-1)
	}
	return a ^ ((a ^ b) & bit)
}
