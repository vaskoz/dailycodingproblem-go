package day69

import "sort"

// LargestProductOf3 returns the largest possible number produced by
// multiplying any three integers.
func LargestProductOf3(nums []int) int {
	smallest := make([]int, 2)
	largest := make([]int, 3)
	copy(smallest, nums[:2])
	copy(largest, nums[:3])
	sort.Ints(smallest)
	sort.Ints(largest)
	if nums[2] < smallest[1] {
		smallest[1] = nums[2]
		if smallest[1] < smallest[0] {
			smallest[0], smallest[1] = smallest[1], smallest[0]
		}
	}

	for i := 3; i < len(nums); i++ {
		num := nums[i]
		if num < smallest[1] {
			smallest[1] = num
			if smallest[1] < smallest[0] {
				smallest[0], smallest[1] = smallest[1], smallest[0]
			}
		}
		if num > largest[0] {
			largest[0] = num
			for i := 1; i < len(largest); i++ {
				if largest[i-1] > largest[i] {
					largest[i-1], largest[i] = largest[i], largest[i-1]
				}
			}
		}
	}

	optionA := smallest[0] * smallest[1] * largest[2]
	optionB := largest[0] * largest[1] * largest[2]
	if optionA > optionB {
		return optionA
	}
	return optionB
}
