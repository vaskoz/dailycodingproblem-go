package day12

import "sort"

// UniqueClimbs calculates the number of possible ways to climb the steps.
// Assumes that we have the given number of possible strides.
// A stride is the number of steps possible.
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

// UniqueClimbsDS doesn't use recursion and instead iterates using memoization of subproblem results.
// Time complexity is O(K*N) where K is the number of possible strides and N is the number of steps.
func UniqueClimbsDS(steps int, strides []int) int {
	copyStrides := make([]int, len(strides))
	copy(copyStrides, strides)
	sort.Ints(copyStrides)
	stepData := make([]int, steps+1)
	stepData[0] = 1
	for step := 1; step < len(stepData); step++ {
		sum := 0
		for _, stride := range strides {
			if step < stride {
				continue
			}
			sum += stepData[step-stride]
		}
		stepData[step] = sum
	}
	return stepData[steps]
}
