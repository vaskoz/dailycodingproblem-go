package day44

// InversionCountBrute returns the inversion count using brute-force.
// Runtime is O(N^2) and O(1) space.
func InversionCountBrute(array []int) int {
	var inversions int
	for i := range array {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				inversions++
			}
		}
	}
	return inversions
}

// InversionCount returns the inversion count.
// Runtime is O(N log N).
func InversionCount(array []int) int {
	return 0
}
