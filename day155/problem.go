package day155

import (
	"errors"
	"sort"
)

var errNoMajority = errors.New("no majority exists")

// ErrNoMajority returns the same error returned when no majority exists.
func ErrNoMajority() error {
	return errNoMajority
}

// MajorityElementMap returns the number that occurs over half the time.
// Runs in O(N) time and O(unique values) in space.
func MajorityElementMap(nums []int) (int, error) {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}
	threshold := len(nums) / 2
	for k, v := range counts {
		if v >= threshold {
			return k, nil
		}
	}
	return 0, errNoMajority
}

// MajorityElementSort returns the number that occurs over half the time.
// Runs in O(N log N) time and O(1) in space.
func MajorityElementSort(nums []int) (int, error) {
	if len(nums) < 1 {
		return 0, errNoMajority
	}
	nums = append([]int{}, nums...) // don't mutate the input.
	sort.Ints(nums)
	var maxVal, maxCount int
	val := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == val {
			count++
		} else {
			if count > maxCount {
				maxCount = count
				maxVal = val
			}
			val = nums[i]
			count = 1
		}
	}
	if count > maxCount {
		return val, nil
	}
	return maxVal, nil
}

// MajorityBoyerMoore runs in O(N) time and O(1) space.
// Limitation: only works if there is definitely a majority value.
func MajorityBoyerMoore(nums []int) (int, error) {
	var m, i int
	for _, x := range nums {
		if i == 0 {
			m = x
			i = 1
		} else if m == x {
			i++
		} else {
			i--
		}
	}
	return m, nil
}
