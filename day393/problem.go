package day393

import "sort"

// LargestRangeBrute returns the largest range of integers
// using brute force.
// Runtime is O(N*K) where K is the length of the longest range.
// Effectively, this is O(N^2).
// Converts the input into a hash set.
func LargestRangeBrute(nums []int) (int, int) {
	if n := len(nums); n == 0 {
		return 0, 0
	} else if n == 1 {
		return nums[0], nums[0]
	}

	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		set[num] = struct{}{}
	}

	var longest, sAns, eAns int

	for start := range set {
		count := 1
		ptr := start + 1

		for _, found := set[ptr]; found; _, found = set[ptr] {
			count++
			ptr++
		}

		if count > longest || (count == longest && start < sAns) {
			longest = count
			sAns = start
			eAns = ptr - 1
		}
	}

	return sAns, eAns
}

// LargestRangeSort returns the largest range of integers
// using slightly better than brute-force.
// Runtime is O(N log N).
// Performs a copy to avoid mutating the input.
func LargestRangeSort(nums []int) (int, int) {
	if n := len(nums); n == 0 {
		return 0, 0
	} else if n == 1 {
		return nums[0], nums[0]
	}

	copied := append([]int{}, nums...)

	sort.Ints(copied)

	var longest, sAns, eAns, start, end, count int

	count = 1
	start = copied[0]

	for i := 1; i < len(copied); i++ {
		if copied[i] == copied[i-1]+1 {
			count++
		} else {
			end = copied[i-1]
			if end-start > longest {
				sAns, eAns = start, end
				longest = count
			}
			start = copied[i]
			count = 1
		}
	}

	end = copied[len(copied)-1]
	if end-start > longest {
		sAns, eAns = start, end
	}

	return sAns, eAns
}
