package day390

// MissingThousandInMillion returns the numbers missing
// between 1 and 1 million.
// Runs in O(N) time and space.
func MissingThousandInMillion(nums []int) []int {
	seen := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		seen[n] = struct{}{}
	}

	ans := make([]int, 0, 1_000_000-len(seen)+1)

	for n := 1; n <= 1_000_000; n++ {
		if _, ok := seen[n]; !ok {
			ans = append(ans, n)
		}
	}

	return ans
}
