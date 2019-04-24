package day245

// MinimumJumps returns the minimum steps to reach the end.
// Each element contains the max steps that can be taken
// in a single jump.
func MinimumJumps(maxSteps []int) int {
	if len(maxSteps) == 1 {
		return 0
	}
	minJumps := int(^uint(0) >> 1)
	if maxSteps[0] == 0 { // can't jump from this cell
		return minJumps
	}
	for jump := maxSteps[0]; jump > 0; jump-- {
		if jump < len(maxSteps) {
			if result := MinimumJumps(maxSteps[jump:]); result < minJumps {
				minJumps = result
			}
		}
	}
	return minJumps + 1
}
