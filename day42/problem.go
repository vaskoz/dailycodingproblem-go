package day42

// SubsetSumBrute returns the subset by performing a brute-force search.
// The first possible result is return if found, otherwise the search is exhaustive.
// The runtime is exponential.
func SubsetSumBrute(ints []int, k int) []int {
	for i, v := range ints {
		if v == k {
			return []int{v}
		} else if k > v {
			if subset := SubsetSumBrute(ints[i+1:], k-v); len(subset) != 0 {
				return append([]int{v}, subset...)
			}
		}
	}
	return nil
}

// SubsetSumDP returns the subset using dynamic programming.
// The runtime is O(N*K) with O(N*K) space.
func SubsetSumDP(ints []int, k int) []int {
	subsets := make([][][]int, len(ints)+1)

	for i := 0; i <= len(ints); i++ {
		subsets[i] = make([][]int, k+1)
		subsets[i][0] = []int{}
	}

	for i := 1; i <= k; i++ {
		subsets[0][i] = nil
	}

	for i := 1; i <= len(ints); i++ {
		for j := 1; j <= k; j++ {
			subsets[i][j] = subsets[i-1][j]
			if j >= ints[i-1] && subsets[i-1][j-ints[i-1]] != nil {
				subsets[i][j] = append([]int{ints[i-1]}, subsets[i-1][j-ints[i-1]]...)
			}
		}
	}

	return subsets[len(ints)][k]
}
