package day40

// FindOnce returns the number that appears only once in a list where
// other numbers appear 3 times.
// Runs in O(N) time and O(1) space.
func FindOnce(nums []int) int {
	var ones, twos int
	for _, v := range nums {
		twos |= ones & v
		ones ^= v
		not := ^(ones & twos)
		ones &= not
		twos &= not
	}
	return ones
}

// FindOnceMap uses a map to find the non-duplicated value.
// Runs in O(N) and O(N) space
func FindOnceMap(nums []int) int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}
	var result int
	for k, v := range counts {
		if v == 1 {
			result = k
			break
		}
	}
	return result
}
