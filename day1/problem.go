package day1

import "sort"

// TwoSumBrute is a brute force implementation.
// O(N^2) time where N is the size of numbers.
// O(1) space
func TwoSumBrute(numbers []int, k int) bool {
	for i, x := range numbers {
		for j, y := range numbers {
			if i != j && x+y == k {
				return true
			}
		}
	}
	return false
}

// TwoSumBetter is asymptotically better than brute force
// O(N log N) time because sorting.
// O(N) space to prevent modifying the original number slice.
func TwoSumBetter(numbers []int, k int) bool {
	copied := append([]int(nil), numbers...)
	sort.Ints(copied)
	for i, x := range copied {
		j := sort.SearchInts(copied, k-x)
		if i != j && j != len(copied) && copied[j] == k-x {
			return true
		}
	}
	return false
}

// TwoSumBest requires only a single pass in the worst case.
// O(N) time
// O(N) space
func TwoSumBest(numbers []int, k int) bool {
	seen := make(map[int]struct{})
	for _, v := range numbers {
		if _, found := seen[k-v]; found {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
