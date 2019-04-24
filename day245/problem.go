package day245

import "errors"

var errImpassable = errors.New("impossible to reach the end")

// ErrImpassable returns the error associated with an impassable
// route.
func ErrImpassable() error {
	return errImpassable
}

// MinimumJumps returns the minimum steps to reach the end.
// Each element contains the max steps that can be taken
// in a single jump.
// Returns an error if impassable.
func MinimumJumps(maxSteps []int) (int, error) {
	if len(maxSteps) < 2 {
		return 0, nil
	}
	minJumps := int(^uint(0) >> 1)
	if maxSteps[0] == 0 { // can't jump from this cell
		return 0, errImpassable
	}
	impassable := true
	for jump := maxSteps[0]; jump > 0; jump-- {
		if jump < len(maxSteps) {
			if result, impass := MinimumJumps(maxSteps[jump:]); impass == nil && result < minJumps {
				impassable = false
				minJumps = result
			}
		}
	}
	if impassable {
		return 0, errImpassable
	}
	return minJumps + 1, nil
}
