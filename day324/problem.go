package day324

import "sort"

// MostMouseSteps returns the most steps taken by the
// mouse that needs to move the farthest.
// Expects len(mice) == len(holes), otherwise panic.
// Runs in O(N log N) time due to sorting both inputs.
func MostMouseSteps(mice, holes []int) int {
	if len(mice) != len(holes) {
		panic("mice and holes lengths do not match")
	}
	sort.Ints(mice)
	sort.Ints(holes)
	n := len(mice)
	var max int
	for i := 0; i < n; i++ {
		move := mice[i] - holes[i]
		if move < 0 {
			move = -move
		}
		if move > max {
			max = move
		}
	}
	return max
}
