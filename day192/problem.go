package day192

// CanAdvanceToEndBrute tries to brute force the answer.
func CanAdvanceToEndBrute(steps []int) bool {
	if len(steps) < 2 {
		return true
	}
	for step := 1; step <= steps[0] && step <= len(steps); step++ {
		if CanAdvanceToEndBrute(steps[step:]) {
			return true
		}
	}
	return false
}
