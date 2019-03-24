package day214

// LengthLongestConsecutiveRunOnesBitShift calculates the longest
// run of 1's using bit shifting.
func LengthLongestConsecutiveRunOnesBitShift(n int) int {
	var max, count int
	for n != 0 {
		if n&1 == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
		n >>= 1
	}
	return max
}
