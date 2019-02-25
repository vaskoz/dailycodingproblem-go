package day186

import "log"

// SubsetsSmallestDifference returns the two subsets with the
// smallest difference between them.
// Brute-force. Runs in O(2^N) time.
func SubsetsSmallestDifference(nums []int) ([]int, []int) {
	perms := make(chan [][]int)
	go func() {
		permutations(nums, []int{}, []int{}, perms)
		close(perms)
	}()
	minSum := int(^uint(0) >> 1)
	var s1, s2 []int
	for subsets := range perms {
		log.Println(subsets)
		if s := sum(subsets[0], subsets[1]); s < minSum {
			minSum = s
			s1 = subsets[0]
			s2 = subsets[1]
		}
	}
	return s1, s2
}

func sum(s1, s2 []int) int {
	sum := 0
	for _, v := range s1 {
		sum += v
	}
	for _, v := range s2 {
		sum -= v
	}
	if sum < 0 {
		return -sum
	}
	return sum
}

func permutations(nums, sub1, sub2 []int, perms chan [][]int) {
	if len(nums) == 0 {
		newS1 := append([]int{}, sub1...)
		newS2 := append([]int{}, sub2...)
		perms <- [][]int{newS1, newS2}
		return
	}
	permutations(nums[1:], append(sub1, nums[0]), sub2, perms)
	permutations(nums[1:], sub1, append(sub2, nums[0]), perms)
}
