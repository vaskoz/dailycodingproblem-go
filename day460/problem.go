package day460

// FlipsXsBeforeYs returns the minimum number of flips
// required to meet the condition that the result should
// have all X's before Y's.
// Runs in O(N) time and O(1) space.
func FlipsXsBeforeYs(input string) int {
	streakX, streakY := true, true

	var changeXtoY, changeYtoX int

	for forward, backward := 0, len(input)-1; forward < len(input); forward, backward = forward+1, backward-1 {
		if streakX && input[forward] != 'x' {
			streakX = false
		}

		if streakY && input[backward] != 'y' {
			streakY = false
		}

		if !streakX && input[forward] == 'x' {
			changeXtoY++
		}

		if !streakY && input[backward] == 'y' {
			changeYtoX++
		}
	}

	if changeXtoY < changeYtoX {
		return changeXtoY
	}

	return changeYtoX
}
