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

// CanAdvanceToEndLinear does it in one pass.
func CanAdvanceToEndLinear(steps []int) bool {
	if len(steps) < 2 {
		return true
	}
	maxReach, step := steps[0], steps[0]
	for i := 1; i < len(steps); i++ {
		if i == len(steps)-1 {
			return true
		}
		maxReach = max(maxReach, i+steps[i])
		step--
		if step == 0 {
			if i >= maxReach {
				break
			}
			step = maxReach - i
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
