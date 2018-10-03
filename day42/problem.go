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
