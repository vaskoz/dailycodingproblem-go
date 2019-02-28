package day189

// LengthLongestDistinct returns the length of the longest
// sub-array where all elements are distinct.
// Runs in O(N) time and O(N) extra space.
func LengthLongestDistinct(nums []int) int {
	valueToPos := make(map[int]int)
	var longest int
	for i, v := range nums {
		if pos, found := valueToPos[v]; found {
			if l := len(valueToPos); l > longest {
				longest = l
			}
			for val, p := range valueToPos {
				if p <= pos {
					delete(valueToPos, val)
				}
			}
		}
		valueToPos[v] = i

	}
	if l := len(valueToPos); l > longest {
		longest = l
	}
	return longest
}
