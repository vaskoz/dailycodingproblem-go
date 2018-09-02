package day12

// UniqueClimbs calculates the number of possible ways to climb the steps.
// Assumes that we have the given number of possible strides.
// A stride is the number of steps possible.
// Assumes strides are in sorted order.
// Runs in exponential time.
func UniqueClimbs(steps int, strides []int) int {
	if steps < 0 {
		return 0
	} else if steps == 0 {
		return 1
	}
	sum := 0
	for _, stride := range strides {
		sum += UniqueClimbs(steps-stride, strides)
	}
	return sum
}
